package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
  "strings"
)

func projectListCommand(c *cli.Context) error {
  client := authorizer.SimpleClient()
  groups, err := client.GetAllProjectsGroupedByCompany()

	if err != nil {
		log.Fatal(err)
	}

	first := true
	for _, group := range groups {
		if first {
			first = false
		} else {
			fmt.Printf("\n")
		}


		fmt.Printf("Company: %s\n", group.CompanyName)

    for _, project := range group.Projects {
      fmt.Printf("- %s (#%v)\n", project.Title, project.ProjectID)
      if project.Notes != "" {
        fmt.Printf("  Notes:\n%s",indent(project.Notes, "    "))
      }
    }
	}

	return nil
}

func indent(what, indentWith string) string {
  var builder  strings.Builder

  for _, part := range strings.Split(what, "\n") {
    builder.WriteString(indentWith)
    builder.WriteString(part)
    builder.WriteString("\n")
  }

  return builder.String()
}
