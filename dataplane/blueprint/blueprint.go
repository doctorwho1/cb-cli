package blueprint

import (
	"strconv"
	"time"

	"github.com/hortonworks/cb-cli/dataplane/oauth"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	v4bp "github.com/hortonworks/cb-cli/dataplane/api/client/v4_workspace_id_blueprints"
	"github.com/hortonworks/cb-cli/dataplane/api/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/dp-cli-common/utils"
	"github.com/urfave/cli"
)

var blueprintHeader = []string{"Name", "Description", "StackName", "StackVersion", "Hostgroup Count", "Tags"}

type blueprintOut struct {
	Name           string `json:"Name" yaml:"Name"`
	Description    string `json:"Description" yaml:"Description"`
	StackName      string `json:"StackName" yaml:"StackName"`
	StackVersion   string `json:"StackVersion" yaml:"StackVersion"`
	HostgroupCount string `json:"HostgroupCount" yaml:"HostgroupCount"`
	Tags           string `json:"Tags" yaml:"Tags"`
}

type blueprintOutJsonDescribe struct {
	*blueprintOut
	Content string `json:"BlueprintTextAsBase64" yaml:"BlueprintTextAsBase64"`
	ID      string `json:"ID" yaml:"ID"`
}

type blueprintOutTableDescribe struct {
	*blueprintOut
	ID string `json:"ID" yaml:"ID"`
}

func (b *blueprintOut) DataAsStringArray() []string {
	return []string{b.Name, b.Description, b.StackName, b.StackVersion, b.HostgroupCount, b.Tags}
}

func (b *blueprintOutJsonDescribe) DataAsStringArray() []string {
	return append(b.blueprintOut.DataAsStringArray(), b.ID)
}

func (b *blueprintOutTableDescribe) DataAsStringArray() []string {
	return append(b.blueprintOut.DataAsStringArray(), b.ID)
}

func CreateBlueprintFromUrl(c *cli.Context) {

	log.Infof("[CreateBlueprintFromUrl] creating blueprint from a URL")
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	urlLocation := c.String(fl.FlURL.Name)
	createBlueprintImpl(
		cbClient.Cloudbreak.V4WorkspaceIDBlueprints,
		c.String(fl.FlName.Name),
		c.String(fl.FlDescriptionOptional.Name),
		c.Bool(fl.FlDlOptional.Name),
		utils.ReadContentFromURL(urlLocation, new(http.Client)),
		c.Int64(fl.FlWorkspaceOptional.Name))
}

func CreateBlueprintFromFile(c *cli.Context) {

	log.Infof("[CreateBlueprintFromFile] creating blueprint from a file")
	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	fileLocation := c.String(fl.FlFile.Name)
	createBlueprintImpl(
		cbClient.Cloudbreak.V4WorkspaceIDBlueprints,
		c.String(fl.FlName.Name),
		c.String(fl.FlDescriptionOptional.Name),
		c.Bool(fl.FlDlOptional.Name),
		utils.ReadFile(fileLocation),
		c.Int64(fl.FlWorkspaceOptional.Name))
}

func createBlueprintImpl(client blueprintClient, name string, description string, dl bool, ambariBlueprint []byte, workspace int64) *model.BlueprintV4Response {
	defer utils.TimeTrack(time.Now(), "create blueprint")
	tags := map[string]interface{}{"shared_services_ready": dl}
	bpRequest := &model.BlueprintV4Request{
		Name:            &name,
		Description:     &description,
		AmbariBlueprint: base64.StdEncoding.EncodeToString(ambariBlueprint),
		Tags:            tags,
	}
	var blueprint *model.BlueprintV4Response
	log.Infof("[createBlueprintImpl] sending create blueprint request")
	resp, err := client.CreateBlueprintInWorkspace(v4bp.NewCreateBlueprintInWorkspaceParams().WithWorkspaceID(workspace).WithBody(bpRequest))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	blueprint = resp.Payload

	log.Infof("[createBlueprintImpl] blueprint created: %s (id: %d)", *blueprint.Name, blueprint.ID)
	return blueprint
}

func DescribeBlueprint(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "describe blueprint")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	bp := FetchBlueprint(c.Int64(fl.FlWorkspaceOptional.Name), c.String(fl.FlName.Name), cbClient.Cloudbreak.V4WorkspaceIDBlueprints)
	if output.Format != "table" {
		output.Write(append(blueprintHeader, "Content", "ID"), convertResponseWithContentAndIDToBlueprint(bp))
	} else {
		output.Write(append(blueprintHeader, "ID"), convertResponseWithIDToBlueprint(bp))
	}
}

type GetBlueprintInWorkspace interface {
	GetBlueprintInWorkspace(*v4bp.GetBlueprintInWorkspaceParams) (*v4bp.GetBlueprintInWorkspaceOK, error)
}

func FetchBlueprint(workspace int64, name string, client GetBlueprintInWorkspace) *model.BlueprintV4Response {
	resp, err := client.GetBlueprintInWorkspace(v4bp.NewGetBlueprintInWorkspaceParams().WithWorkspaceID(workspace).WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	return resp.Payload
}

func DeleteBlueprint(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "delete blueprint")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	deleteBlueprintsImpl(cbClient.Cloudbreak.V4WorkspaceIDBlueprints, c.Int64(fl.FlWorkspaceOptional.Name), c.String(fl.FlName.Name))
}

func deleteBlueprintsImpl(client blueprintClient, workspace int64, name string) {
	log.Infof("[deleteBlueprintsImpl] sending delete blueprint request with name: %s", name)
	_, err := client.DeleteBlueprintInWorkspace(v4bp.NewDeleteBlueprintInWorkspaceParams().WithWorkspaceID(workspace).WithName(name))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[deleteBlueprintsImpl] blueprint deleted, name: %s", name)
}

func ListBlueprints(c *cli.Context) {
	defer utils.TimeTrack(time.Now(), "get blueprints")

	cbClient := oauth.NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	workspace := fl.FlWorkspaceOptional.Name
	listBlueprintsImpl(c.Int64(workspace), cbClient.Cloudbreak.V4WorkspaceIDBlueprints, output.WriteList)
}

type blueprintClient interface {
	CreateBlueprintInWorkspace(params *v4bp.CreateBlueprintInWorkspaceParams) (*v4bp.CreateBlueprintInWorkspaceOK, error)
	ListBlueprintsByWorkspace(params *v4bp.ListBlueprintsByWorkspaceParams) (*v4bp.ListBlueprintsByWorkspaceOK, error)
	DeleteBlueprintInWorkspace(params *v4bp.DeleteBlueprintInWorkspaceParams) (*v4bp.DeleteBlueprintInWorkspaceOK, error)
}

func listBlueprintsImpl(workspace int64, client blueprintClient, writer func([]string, []utils.Row)) {
	log.Infof("[listBlueprintsImpl] sending blueprint list request")
	resp, err := client.ListBlueprintsByWorkspace(v4bp.NewListBlueprintsByWorkspaceParams().WithWorkspaceID(workspace))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	var tableRows []utils.Row
	for _, bp := range resp.Payload.Responses {
		tableRows = append(tableRows, convertResponseToBlueprint(bp))
	}

	writer(blueprintHeader, tableRows)
}

func convertResponseToBlueprint(bp *model.BlueprintV4ViewResponse) *blueprintOut {
	return &blueprintOut{
		Name:           *bp.Name,
		Description:    utils.SafeStringConvert(bp.Description),
		StackName:      fmt.Sprintf("%v", bp.StackType),
		StackVersion:   fmt.Sprintf("%v", bp.StackVersion),
		HostgroupCount: fmt.Sprint(bp.HostGroupCount),
		Tags:           bp.Status,
	}
}

func convertResponseWithContentAndIDToBlueprint(bp *model.BlueprintV4Response) *blueprintOutJsonDescribe {
	jsonRoot := decodeAndParseToJson(bp.AmbariBlueprint)
	blueprintsNode := jsonRoot["Blueprints"].(map[string]interface{})
	return &blueprintOutJsonDescribe{
		blueprintOut: &blueprintOut{
			Name:           *bp.Name,
			Description:    *bp.Description,
			StackName:      fmt.Sprintf("%v", blueprintsNode["stack_name"]),
			StackVersion:   fmt.Sprintf("%v", blueprintsNode["stack_version"]),
			HostgroupCount: fmt.Sprint(bp.HostGroupCount),
			Tags:           bp.Status,
		},
		Content: bp.AmbariBlueprint,
		ID:      strconv.FormatInt(bp.ID, 10),
	}
}

func convertResponseWithIDToBlueprint(bp *model.BlueprintV4Response) *blueprintOutTableDescribe {
	jsonRoot := decodeAndParseToJson(bp.AmbariBlueprint)
	blueprintsNode := jsonRoot["Blueprints"].(map[string]interface{})
	return &blueprintOutTableDescribe{
		blueprintOut: &blueprintOut{
			Name:           *bp.Name,
			Description:    *bp.Description,
			StackName:      fmt.Sprintf("%v", blueprintsNode["stack_name"]),
			StackVersion:   fmt.Sprintf("%v", blueprintsNode["stack_version"]),
			HostgroupCount: fmt.Sprint(bp.HostGroupCount),
			Tags:           bp.Status,
		},
		ID: strconv.FormatInt(bp.ID, 10),
	}
}

func decodeAndParseToJson(encodedBlueprint string) map[string]interface{} {
	log.Debugf("[decodeAndParseToJson] decoding blueprint from base64")
	b, err := base64.StdEncoding.DecodeString(encodedBlueprint)
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Debugf("[decodeAndParseToJson] parse blueprint to JSON")
	var blueprintJson map[string]interface{}
	if err = json.Unmarshal(b, &blueprintJson); err != nil {
		utils.LogErrorAndExit(err)
	}
	return blueprintJson
}
