package utils

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/ory/dockertest/v3"
	"log"
)

var DbTest *sql.DB

func IntegrationTestSetup() (*dockertest.Pool, *dockertest.Resource) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mariadb", "latest", []string{"MYSQL_ROOT_PASSWORD=password-test"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		DbTest, err = sql.Open("mysql", fmt.Sprintf("root:password-test@(localhost:%s)/mysql?multiStatements=true", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return DbTest.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	DbTest.SetMaxIdleConns(0)
	driver, _ := mysql.WithInstance(DbTest, &mysql.Config{})
	migration, err := migrate.NewWithDatabaseInstance("file://../database/migrations/", "mysql", driver)
	if err != nil {
		log.Fatalf("Error running migrations: %s", err)
	}
	err = migration.Up()
	if err != nil {
		log.Fatal(err.Error())
	}

	return pool, resource
}

func IntegrationTestTeardown(pool *dockertest.Pool, resource *dockertest.Resource) {
	if err := pool.Purge(resource); err != nil {
		fmt.Printf("Could not purge resource: %s\n", err)
	}
}
