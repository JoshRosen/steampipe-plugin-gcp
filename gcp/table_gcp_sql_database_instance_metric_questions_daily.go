package gcp

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

//// TABLE DEFINITION

func tableGcpSQLDatabaseInstanceMetricQuestionsDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "gcp_sql_database_instance_metric_questions_daily",
		Description: "GCP SQL Database Instance Metrics - Questions (Daily)",
		List: &plugin.ListConfig{
			ParentHydrate:     listSQLDatabaseInstances,
			Hydrate:           listSQLDatabaseInstanceMetricQuestionsDaily,
			ShouldIgnoreError: isIgnorableError([]string{"403"}),
		},
		Columns: monitoringMetricColumns([]*plugin.Column{
			{
				Name:        "instance_id",
				Description: "The ID of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DimensionValue"),
			},
		}),
	}
}

//// LIST FUNCTION

func listSQLDatabaseInstanceMetricQuestionsDaily(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	instanceInfo := h.Item.(*sqladmin.DatabaseInstance)

	// Get project details
	getProjectCached := plugin.HydrateFunc(getProject).WithCache()
	projectId, err := getProjectCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	project := projectId.(string)
	location := instanceInfo.Region
	dimensionValue := "\"" + project + ":" + instanceInfo.Name + "\""

	return listMonitorMetricStatistics(ctx, d, h, "DAILY", "\"cloudsql.googleapis.com/database/mysql/questions\"", "resource.label.database_id = ", dimensionValue, instanceInfo.Name, location)
}
