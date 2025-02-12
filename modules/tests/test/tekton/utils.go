package tekton

import (
	"fmt"
	"strings"

	beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

func TaskResultsToMap(results []beta1.TaskRunResult) map[string]string {
	mappedResult := make(map[string]string, len(results))

	for _, result := range results {
		mappedResult[result.Name] = resultToString(result.Value)
	}

	return mappedResult
}

func PipelineResultsToMap(results []beta1.PipelineRunResult) map[string]string {
	mappedResult := make(map[string]string, len(results))

	for _, result := range results {
		mappedResult[result.Name] = resultToString(result.Value)
	}

	return mappedResult
}

func resultToString(value beta1.ParamValue) string {
	if value.StringVal != "" {
		return value.StringVal
	}
	if len(value.ObjectVal) > 0 {
		return objectToString(value.ObjectVal)
	}

	return arrayToString(value.ArrayVal)
}

func arrayToString(value []string) string {
	return strings.Join(value, ", ")
}

func objectToString(value map[string]string) string {
	res := make([]string, len(value))
	i := 0
	for key, val := range value {
		res[i] = fmt.Sprintf("%s - %s", key, val)
		i++
	}
	return strings.Join(res, ", ")
}
