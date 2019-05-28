package migrate

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AssessmentsClient is the move your workloads to Azure.
type AssessmentsClient struct {
	BaseClient
}

// NewAssessmentsClient creates an instance of the AssessmentsClient client.
func NewAssessmentsClient(subscriptionID string, acceptLanguage string) AssessmentsClient {
	return NewAssessmentsClientWithBaseURI(DefaultBaseURI, subscriptionID, acceptLanguage)
}

// NewAssessmentsClientWithBaseURI creates an instance of the AssessmentsClient client.
func NewAssessmentsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) AssessmentsClient {
	return AssessmentsClient{NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)}
}

// Create create a new assessment with the given name and the specified settings. Since name of an assessment in a
// project is a unique identifier, if an assessment with the name provided already exists, then the existing assessment
// is updated.
//
// Any PUT operation, resulting in either create or update on an assessment, will cause the assessment to go in a
// "InProgress" state. This will be indicated by the field 'computationState' on the Assessment object. During this
// time no other PUT operation will be allowed on that assessment object, nor will a Delete operation. Once the
// computation for the assessment is complete, the field 'computationState' will be updated to 'Ready', and then other
// PUT or DELETE operations can happen on the assessment.
//
// When assessment is under computation, any PUT will lead to a 400 - Bad Request error.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// projectName - name of the Azure Migrate project.
// groupName - unique name of a group within a project.
// assessmentName - unique name of an assessment within a project.
// assessment - new or Updated Assessment object.
func (client AssessmentsClient) Create(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessment *Assessment) (result Assessment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: assessment,
			Constraints: []validation.Constraint{{Target: "assessment", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "assessment.AssessmentProperties", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "assessment.AssessmentProperties.ScalingFactor", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "assessment.AssessmentProperties.DiscountPercentage", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("migrate.AssessmentsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, projectName, groupName, assessmentName, assessment)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AssessmentsClient) CreatePreparer(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessment *Assessment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assessmentName":    autorest.Encode("path", assessmentName),
		"groupName":         autorest.Encode("path", groupName),
		"projectName":       autorest.Encode("path", projectName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	assessment.ID = nil
	assessment.Name = nil
	assessment.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if assessment != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(assessment))
	}
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AssessmentsClient) CreateResponder(resp *http.Response) (result Assessment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an assessment from the project. The machines remain in the assessment. Deleting a non-existent
// assessment results in a no-operation.
//
// When an assessment is under computation, as indicated by the 'computationState' field, it cannot be deleted. Any
// such attempt will return a 400 - Bad Request.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// projectName - name of the Azure Migrate project.
// groupName - unique name of a group within a project.
// assessmentName - unique name of an assessment within a project.
func (client AssessmentsClient) Delete(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, projectName, groupName, assessmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AssessmentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assessmentName":    autorest.Encode("path", assessmentName),
		"groupName":         autorest.Encode("path", groupName),
		"projectName":       autorest.Encode("path", projectName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AssessmentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get an existing assessment with the specified name. Returns a json object of type 'assessment' as specified in
// Models section.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// projectName - name of the Azure Migrate project.
// groupName - unique name of a group within a project.
// assessmentName - unique name of an assessment within a project.
func (client AssessmentsClient) Get(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string) (result Assessment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, projectName, groupName, assessmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AssessmentsClient) GetPreparer(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assessmentName":    autorest.Encode("path", assessmentName),
		"groupName":         autorest.Encode("path", groupName),
		"projectName":       autorest.Encode("path", projectName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssessmentsClient) GetResponder(resp *http.Response) (result Assessment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetReportDownloadURL get the URL for downloading the assessment in a report format.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// projectName - name of the Azure Migrate project.
// groupName - unique name of a group within a project.
// assessmentName - unique name of an assessment within a project.
func (client AssessmentsClient) GetReportDownloadURL(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string) (result DownloadURL, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsClient.GetReportDownloadURL")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetReportDownloadURLPreparer(ctx, resourceGroupName, projectName, groupName, assessmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "GetReportDownloadURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetReportDownloadURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "GetReportDownloadURL", resp, "Failure sending request")
		return
	}

	result, err = client.GetReportDownloadURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "GetReportDownloadURL", resp, "Failure responding to request")
	}

	return
}

// GetReportDownloadURLPreparer prepares the GetReportDownloadURL request.
func (client AssessmentsClient) GetReportDownloadURLPreparer(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assessmentName":    autorest.Encode("path", assessmentName),
		"groupName":         autorest.Encode("path", groupName),
		"projectName":       autorest.Encode("path", projectName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}/downloadUrl", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetReportDownloadURLSender sends the GetReportDownloadURL request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentsClient) GetReportDownloadURLSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetReportDownloadURLResponder handles the response to the GetReportDownloadURL request. The method always
// closes the http.Response Body.
func (client AssessmentsClient) GetReportDownloadURLResponder(resp *http.Response) (result DownloadURL, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByGroup get all assessments created for the specified group.
//
// Returns a json array of objects of type 'assessment' as specified in Models section.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// projectName - name of the Azure Migrate project.
// groupName - unique name of a group within a project.
func (client AssessmentsClient) ListByGroup(ctx context.Context, resourceGroupName string, projectName string, groupName string) (result AssessmentResultList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsClient.ListByGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByGroupPreparer(ctx, resourceGroupName, projectName, groupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "ListByGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "ListByGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "ListByGroup", resp, "Failure responding to request")
	}

	return
}

// ListByGroupPreparer prepares the ListByGroup request.
func (client AssessmentsClient) ListByGroupPreparer(ctx context.Context, resourceGroupName string, projectName string, groupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupName":         autorest.Encode("path", groupName),
		"projectName":       autorest.Encode("path", projectName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByGroupSender sends the ListByGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentsClient) ListByGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByGroupResponder handles the response to the ListByGroup request. The method always
// closes the http.Response Body.
func (client AssessmentsClient) ListByGroupResponder(resp *http.Response) (result AssessmentResultList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByProject get all assessments created in the project.
//
// Returns a json array of objects of type 'assessment' as specified in Models section.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// projectName - name of the Azure Migrate project.
func (client AssessmentsClient) ListByProject(ctx context.Context, resourceGroupName string, projectName string) (result AssessmentResultList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentsClient.ListByProject")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByProjectPreparer(ctx, resourceGroupName, projectName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "ListByProject", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByProjectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "ListByProject", resp, "Failure sending request")
		return
	}

	result, err = client.ListByProjectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentsClient", "ListByProject", resp, "Failure responding to request")
	}

	return
}

// ListByProjectPreparer prepares the ListByProject request.
func (client AssessmentsClient) ListByProjectPreparer(ctx context.Context, resourceGroupName string, projectName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"projectName":       autorest.Encode("path", projectName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/assessments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByProjectSender sends the ListByProject request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentsClient) ListByProjectSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByProjectResponder handles the response to the ListByProject request. The method always
// closes the http.Response Body.
func (client AssessmentsClient) ListByProjectResponder(resp *http.Response) (result AssessmentResultList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
