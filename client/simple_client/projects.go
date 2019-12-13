package simple_client

import (
  "github.com/moretea/avaza/client/models"
  "github.com/moretea/avaza/client/client/project"
  "sort"
)

// Get all projects accessible to the current user, uses pagination.
func (c *Client) GetAllProjects() ([]*models.ProjectListDetails, error) {
	c.authorizer.RefreshIfNecessary()
  var pageNumber int32 = 1
  var pageSize int32 = 1000

  params := project.NewProjectGetParams().
    WithPageNumber(&pageNumber).
    WithPageSize(&pageSize)

  var result []*models.ProjectListDetails;

  for {
    projectOk, err := c.avazaClient.Project.ProjectGet(params, c.authorizer.CreateAuth())

    if err != nil {
      return nil, err
    }

    for _, project := range projectOk.Payload.Projects {
      result = append(result, project)
    }

    if int32(len(result)) >= projectOk.Payload.TotalCount {
      break
    } else {
      pageNumber = pageNumber + 1
    }
  }

  sort.Slice(result, func(i, j int) bool {
    return result[i].Title < result[j].Title
  })

  return result, nil
}

type ProjectsGroupedByCompany struct {
  CompanyIDFK int32
  CompanyName string

  Projects []*models.ProjectListDetails
}

func (c *Client) GetAllProjectsGroupedByCompany() ([]*ProjectsGroupedByCompany, error) {
  all, err := c.GetAllProjects()

  if err != nil {
    return nil, err
  }

  temp := make(map[int32]*ProjectsGroupedByCompany,0)

  for _, project := range all {
    if _, ok := temp[project.CompanyIDFK]; !ok {
      temp[project.CompanyIDFK] = &ProjectsGroupedByCompany {
        CompanyIDFK: project.CompanyIDFK,
        CompanyName: project.CompanyName,
      }
    }

    temp[project.CompanyIDFK].Projects = append(temp[project.CompanyIDFK].Projects, project)
  }

  var result []*ProjectsGroupedByCompany
  for key := range temp {
    result = append(result, temp[key])
  }

  // Sort by Company first
  sort.Slice(result, func(i, j int) bool {
      return result[i].CompanyName < result[j].CompanyName
  })

  // Within each group, sort by project name.


  return result, nil
}
