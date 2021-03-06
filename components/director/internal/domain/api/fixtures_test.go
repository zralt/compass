package api_test

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/kyma-incubator/compass/components/director/pkg/str"

	"github.com/kyma-incubator/compass/components/director/internal/domain/version"

	"github.com/kyma-incubator/compass/components/director/internal/domain/api"
	"github.com/kyma-incubator/compass/components/director/internal/repo"

	"github.com/kyma-incubator/compass/components/director/internal/model"
	"github.com/kyma-incubator/compass/components/director/pkg/graphql"
)

const (
	apiDefID  = "ddddddddd-dddd-dddd-dddd-dddddddddddd"
	appID     = "aaaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
	tenantID  = "ttttttttt-tttt-tttt-tttt-tttttttttttt"
	packageID = "ppppppppp-pppp-pppp-pppp-pppppppppppp"
)

func fixAPIDefinitionModel(id string, appId, pkgID *string, name, targetURL string) *model.APIDefinition {
	return &model.APIDefinition{
		ID:            id,
		ApplicationID: appId,
		PackageID:     pkgID,
		Name:          name,
		TargetURL:     targetURL,
	}
}

func fixFullAPIDefinitionModelWithAPIRtmAuth(placeholder string) *model.APIDefinition {
	apiModel := fixFullAPIDefinitionModel(placeholder)

	apiRtmAuth := model.APIRuntimeAuth{
		ID:        str.Ptr("foo"),
		TenantID:  "tnt",
		RuntimeID: "1",
		APIDefID:  "2",
		Value:     apiModel.DefaultAuth,
	}
	apiModel.Auths = []*model.APIRuntimeAuth{&apiRtmAuth, &apiRtmAuth}
	return &apiModel
}

func fixFullAPIDefinitionModel(placeholder string) model.APIDefinition {
	spec := &model.APISpec{
		Data:   str.Ptr("spec_data_" + placeholder),
		Format: model.SpecFormatYaml,
		Type:   model.APISpecTypeOpenAPI,
	}

	deprecated := false
	forRemoval := false

	v := &model.Version{
		Value:           "v1.1",
		Deprecated:      &deprecated,
		DeprecatedSince: str.Ptr("v1.0"),
		ForRemoval:      &forRemoval,
	}

	auth := model.Auth{
		AdditionalHeaders: map[string][]string{"testHeader": {"hval1", "hval2"}},
	}

	return model.APIDefinition{
		ID:            apiDefID,
		ApplicationID: str.Ptr(appID),
		Tenant:        tenantID,
		PackageID:     str.Ptr(packageID),
		Name:          placeholder,
		Description:   str.Ptr("desc_" + placeholder),
		Spec:          spec,
		TargetURL:     fmt.Sprintf("https://%s.com", placeholder),
		Group:         str.Ptr("group_" + placeholder),
		DefaultAuth:   &auth,
		Version:       v,
	}
}

func fixGQLAPIDefinition(id string, appId, pkgId *string, name, targetURL string) *graphql.APIDefinition {
	return &graphql.APIDefinition{
		ID:            id,
		ApplicationID: appId,
		PackageID:     pkgId,
		Name:          name,
		TargetURL:     targetURL,
	}
}

func fixFullGQLAPIDefinition(placeholder string) *graphql.APIDefinition {
	data := graphql.CLOB("spec_data_" + placeholder)
	format := graphql.SpecFormatYaml

	spec := &graphql.APISpec{
		Data:         &data,
		Format:       format,
		Type:         graphql.APISpecTypeOpenAPI,
		DefinitionID: apiDefID,
	}

	deprecated := false
	forRemoval := false

	v := &graphql.Version{
		Value:           "v1.1",
		Deprecated:      &deprecated,
		DeprecatedSince: str.Ptr("v1.0"),
		ForRemoval:      &forRemoval,
	}

	headers := graphql.HttpHeaders{"testHeader": {"hval1", "hval2"}}

	auth := graphql.Auth{
		AdditionalHeaders: &headers,
	}

	return &graphql.APIDefinition{
		ID:            apiDefID,
		ApplicationID: str.Ptr(appID),
		PackageID:     str.Ptr(packageID),
		Name:          placeholder,
		Description:   str.Ptr("desc_" + placeholder),
		Spec:          spec,
		TargetURL:     fmt.Sprintf("https://%s.com", placeholder),
		Group:         str.Ptr("group_" + placeholder),
		DefaultAuth:   &auth,
		Version:       v,
	}
}

func fixModelAPIDefinitionInput(name, description string, group string) *model.APIDefinitionInput {
	data := "data"

	spec := &model.APISpecInput{
		Data:         &data,
		Type:         model.APISpecTypeOpenAPI,
		Format:       model.SpecFormatYaml,
		FetchRequest: &model.FetchRequestInput{},
	}

	deprecated := false
	deprecatedSince := ""
	forRemoval := false

	v := &model.VersionInput{
		Value:           "1.0.0",
		Deprecated:      &deprecated,
		DeprecatedSince: &deprecatedSince,
		ForRemoval:      &forRemoval,
	}

	basicCredentialDataInput := model.BasicCredentialDataInput{
		Username: "test",
		Password: "pwd",
	}
	authInput := model.AuthInput{
		Credential: &model.CredentialDataInput{Basic: &basicCredentialDataInput},
	}

	return &model.APIDefinitionInput{
		Name:        name,
		Description: &description,
		TargetURL:   "https://test-url.com",
		Group:       &group,
		Spec:        spec,
		Version:     v,
		DefaultAuth: &authInput,
	}
}

func fixGQLAPIDefinitionInput(name, description string, group string) *graphql.APIDefinitionInput {
	data := graphql.CLOB("data")

	spec := &graphql.APISpecInput{
		Data:         &data,
		Type:         graphql.APISpecTypeOpenAPI,
		Format:       graphql.SpecFormatYaml,
		FetchRequest: &graphql.FetchRequestInput{},
	}

	deprecated := false
	deprecatedSince := ""
	forRemoval := false

	v := &graphql.VersionInput{
		Value:           "1.0.0",
		Deprecated:      &deprecated,
		DeprecatedSince: &deprecatedSince,
		ForRemoval:      &forRemoval,
	}

	basicCredentialDataInput := graphql.BasicCredentialDataInput{
		Username: "test",
		Password: "pwd",
	}

	credentialDataInput := graphql.CredentialDataInput{Basic: &basicCredentialDataInput}
	defaultAuth := graphql.AuthInput{
		Credential: &credentialDataInput,
	}

	return &graphql.APIDefinitionInput{
		Name:        name,
		Description: &description,
		TargetURL:   "https://test-url.com",
		Group:       &group,
		Spec:        spec,
		Version:     v,
		DefaultAuth: &defaultAuth,
	}
}

func fixModelAuthInput(headers map[string][]string) *model.AuthInput {
	return &model.AuthInput{
		AdditionalHeaders: headers,
	}
}

func fixGQLAuthInput(headers map[string][]string) *graphql.AuthInput {
	httpHeaders := graphql.HttpHeaders(headers)

	return &graphql.AuthInput{
		AdditionalHeaders: &httpHeaders,
	}
}

func fixModelAuth() *model.Auth {
	return &model.Auth{
		Credential: model.CredentialData{
			Basic: &model.BasicCredentialData{
				Username: "foo",
				Password: "bar",
			},
		},
		AdditionalHeaders:     map[string][]string{"test": {"foo", "bar"}},
		AdditionalQueryParams: map[string][]string{"test": {"foo", "bar"}},
		RequestAuth: &model.CredentialRequestAuth{
			Csrf: &model.CSRFTokenCredentialRequestAuth{
				TokenEndpointURL: "foo.url",
				Credential: model.CredentialData{
					Basic: &model.BasicCredentialData{
						Username: "boo",
						Password: "far",
					},
				},
				AdditionalHeaders:     map[string][]string{"test": {"foo", "bar"}},
				AdditionalQueryParams: map[string][]string{"test": {"foo", "bar"}},
			},
		},
	}
}

func fixGQLAuth() *graphql.Auth {
	return &graphql.Auth{
		Credential: &graphql.BasicCredentialData{
			Username: "foo",
			Password: "bar",
		},
		AdditionalHeaders:     &graphql.HttpHeaders{"test": {"foo", "bar"}},
		AdditionalQueryParams: &graphql.QueryParams{"test": {"foo", "bar"}},
		RequestAuth: &graphql.CredentialRequestAuth{
			Csrf: &graphql.CSRFTokenCredentialRequestAuth{
				TokenEndpointURL: "foo.url",
				Credential: &graphql.BasicCredentialData{
					Username: "boo",
					Password: "far",
				},
				AdditionalHeaders:     &graphql.HttpHeaders{"test": {"foo", "bar"}},
				AdditionalQueryParams: &graphql.QueryParams{"test": {"foo", "bar"}},
			},
		},
	}
}

func fixModelAPIRtmAuth(id string, auth *model.Auth) *model.APIRuntimeAuth {
	return &model.APIRuntimeAuth{
		ID:        str.Ptr("foo"),
		TenantID:  "tnt",
		RuntimeID: id,
		APIDefID:  "api_id",
		Value:     auth,
	}
}

func fixGQLAPIRtmAuth(id string, auth *graphql.Auth) *graphql.APIRuntimeAuth {
	return &graphql.APIRuntimeAuth{
		RuntimeID: id,
		Auth:      auth,
	}
}

func fixEntityAPIDefinition(id string, appId, pkgID *string, name, targetUrl string) api.Entity {
	return api.Entity{
		ID:        id,
		AppID:     repo.NewNullableString(appId),
		PkgID:     repo.NewNullableString(pkgID),
		Name:      name,
		TargetURL: targetUrl,
	}
}

func fixFullEntityAPIDefinition(apiDefID, placeholder string) api.Entity {
	boolPlaceholder := false

	return api.Entity{
		ID:          apiDefID,
		TenantID:    tenantID,
		AppID:       repo.NewNullableString(str.Ptr(appID)),
		PkgID:       repo.NewNullableString(str.Ptr(packageID)),
		Name:        placeholder,
		Description: repo.NewValidNullableString("desc_" + placeholder),
		Group:       repo.NewValidNullableString("group_" + placeholder),
		TargetURL:   fmt.Sprintf("https://%s.com", placeholder),
		EntitySpec: api.EntitySpec{
			SpecData:   repo.NewValidNullableString("spec_data_" + placeholder),
			SpecFormat: repo.NewValidNullableString(string(model.SpecFormatYaml)),
			SpecType:   repo.NewValidNullableString(string(model.APISpecTypeOpenAPI)),
		},
		DefaultAuth: repo.NewValidNullableString(fixDefaultAuth()),
		Version: version.Version{
			VersionValue:           repo.NewNullableString(str.Ptr("v1.1")),
			VersionDepracated:      repo.NewNullableBool(&boolPlaceholder),
			VersionDepracatedSince: repo.NewNullableString(str.Ptr("v1.0")),
			VersionForRemoval:      repo.NewNullableBool(&boolPlaceholder),
		},
	}
}

func fixAPIDefinitionColumns() []string {
	return []string{"id", "tenant_id", "app_id", "package_id", "name", "description", "group_name", "target_url", "spec_data",
		"spec_format", "spec_type", "default_auth", "version_value", "version_deprecated",
		"version_deprecated_since", "version_for_removal"}
}

func fixAPIDefinitionRow(id, placeholder string) []driver.Value {
	return []driver.Value{id, tenantID, appID, packageID, placeholder, "desc_" + placeholder, "group_" + placeholder,
		fmt.Sprintf("https://%s.com", placeholder), "spec_data_" + placeholder, "YAML", "OPEN_API",
		fixDefaultAuth(), "v1.1", false, "v1.0", false}
}

func fixAPICreateArgs(id, defAuth string, api *model.APIDefinition) []driver.Value {
	return []driver.Value{id, tenantID, appID, packageID, api.Name, api.Description, api.Group,
		api.TargetURL, api.Spec.Data, string(api.Spec.Format), string(api.Spec.Type),
		defAuth, api.Version.Value, api.Version.Deprecated, api.Version.DeprecatedSince,
		api.Version.ForRemoval}
}

func fixDefaultAuth() string {
	return `{"Credential":{"Basic":null,"Oauth":null},"AdditionalHeaders":{"testHeader":["hval1","hval2"]},"AdditionalQueryParams":null,"RequestAuth":null}`
}

func fixModelFetchRequest(id, url string, timestamp time.Time) *model.FetchRequest {
	return &model.FetchRequest{
		ID:     id,
		Tenant: tenantID,
		URL:    url,
		Auth:   nil,
		Mode:   "SINGLE",
		Filter: nil,
		Status: &model.FetchRequestStatus{
			Condition: model.FetchRequestStatusConditionInitial,
			Timestamp: timestamp,
		},
		ObjectType: model.APIFetchRequestReference,
		ObjectID:   "foo",
	}
}

func fixGQLFetchRequest(url string, timestamp time.Time) *graphql.FetchRequest {
	return &graphql.FetchRequest{
		Filter: nil,
		Mode:   graphql.FetchModeSingle,
		Auth:   nil,
		URL:    url,
		Status: &graphql.FetchRequestStatus{
			Timestamp: graphql.Timestamp(timestamp),
			Condition: graphql.FetchRequestStatusConditionInitial,
		},
	}
}
