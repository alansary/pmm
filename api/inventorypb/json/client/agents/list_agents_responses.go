// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// ListAgentsReader is a Reader for the ListAgents structure.
type ListAgentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAgentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAgentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListAgentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAgentsOK creates a ListAgentsOK with default headers values
func NewListAgentsOK() *ListAgentsOK {
	return &ListAgentsOK{}
}

/*ListAgentsOK handles this case with default header values.

A successful response.
*/
type ListAgentsOK struct {
	Payload *ListAgentsOKBody
}

func (o *ListAgentsOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/List][%d] listAgentsOk  %+v", 200, o.Payload)
}

func (o *ListAgentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListAgentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAgentsDefault creates a ListAgentsDefault with default headers values
func NewListAgentsDefault(code int) *ListAgentsDefault {
	return &ListAgentsDefault{
		_statusCode: code,
	}
}

/*ListAgentsDefault handles this case with default header values.

An error response.
*/
type ListAgentsDefault struct {
	_statusCode int

	Payload *ListAgentsDefaultBody
}

// Code gets the status code for the list agents default response
func (o *ListAgentsDefault) Code() int {
	return o._statusCode
}

func (o *ListAgentsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/List][%d] ListAgents default  %+v", o._statusCode, o.Payload)
}

func (o *ListAgentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListAgentsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ExternalExporterItems0 ExternalExporter does not run on any Inventory Node.
swagger:model ExternalExporterItems0
*/
type ExternalExporterItems0 struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// URL for scraping metrics.
	MetricsURL string `json:"metrics_url,omitempty"`
}

// Validate validates this external exporter items0
func (o *ExternalExporterItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExternalExporterItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExternalExporterItems0) UnmarshalBinary(b []byte) error {
	var res ExternalExporterItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListAgentsBody list agents body
swagger:model ListAgentsBody
*/
type ListAgentsBody struct {

	// Return only Agents that provide insights for that Node.
	NodeID string `json:"node_id,omitempty"`

	// Return only Agents started by this pmm-agent.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Return only Agents that provide insights for that Service.
	ServiceID string `json:"service_id,omitempty"`
}

// Validate validates this list agents body
func (o *ListAgentsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListAgentsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAgentsBody) UnmarshalBinary(b []byte) error {
	var res ListAgentsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListAgentsDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model ListAgentsDefaultBody
*/
type ListAgentsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this list agents default body
func (o *ListAgentsDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListAgentsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAgentsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListAgentsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListAgentsOKBody list agents OK body
swagger:model ListAgentsOKBody
*/
type ListAgentsOKBody struct {

	// external exporter
	ExternalExporter []*ExternalExporterItems0 `json:"external_exporter"`

	// mongodb exporter
	MongodbExporter []*MongodbExporterItems0 `json:"mongodb_exporter"`

	// mysqld exporter
	MysqldExporter []*MysqldExporterItems0 `json:"mysqld_exporter"`

	// node exporter
	NodeExporter []*NodeExporterItems0 `json:"node_exporter"`

	// pmm agent
	PMMAgent []*PMMAgentItems0 `json:"pmm_agent"`

	// postgres exporter
	PostgresExporter []*PostgresExporterItems0 `json:"postgres_exporter"`

	// qan mongo profiler agent
	QANMongoProfilerAgent []*QANMongoProfilerAgentItems0 `json:"qan_mongo_profiler_agent"`

	// qan mysql perfschema agent
	QANMysqlPerfschemaAgent []*QANMysqlPerfschemaAgentItems0 `json:"qan_mysql_perfschema_agent"`

	// rds exporter
	RDSExporter []*RDSExporterItems0 `json:"rds_exporter"`
}

// Validate validates this list agents OK body
func (o *ListAgentsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExternalExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMongodbExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMysqldExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNodeExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePMMAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePostgresExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQANMongoProfilerAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQANMysqlPerfschemaAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRDSExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListAgentsOKBody) validateExternalExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.ExternalExporter) { // not required
		return nil
	}

	for i := 0; i < len(o.ExternalExporter); i++ {
		if swag.IsZero(o.ExternalExporter[i]) { // not required
			continue
		}

		if o.ExternalExporter[i] != nil {
			if err := o.ExternalExporter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listAgentsOk" + "." + "external_exporter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListAgentsOKBody) validateMongodbExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MongodbExporter) { // not required
		return nil
	}

	for i := 0; i < len(o.MongodbExporter); i++ {
		if swag.IsZero(o.MongodbExporter[i]) { // not required
			continue
		}

		if o.MongodbExporter[i] != nil {
			if err := o.MongodbExporter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listAgentsOk" + "." + "mongodb_exporter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListAgentsOKBody) validateMysqldExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MysqldExporter) { // not required
		return nil
	}

	for i := 0; i < len(o.MysqldExporter); i++ {
		if swag.IsZero(o.MysqldExporter[i]) { // not required
			continue
		}

		if o.MysqldExporter[i] != nil {
			if err := o.MysqldExporter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listAgentsOk" + "." + "mysqld_exporter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListAgentsOKBody) validateNodeExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.NodeExporter) { // not required
		return nil
	}

	for i := 0; i < len(o.NodeExporter); i++ {
		if swag.IsZero(o.NodeExporter[i]) { // not required
			continue
		}

		if o.NodeExporter[i] != nil {
			if err := o.NodeExporter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listAgentsOk" + "." + "node_exporter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListAgentsOKBody) validatePMMAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.PMMAgent) { // not required
		return nil
	}

	for i := 0; i < len(o.PMMAgent); i++ {
		if swag.IsZero(o.PMMAgent[i]) { // not required
			continue
		}

		if o.PMMAgent[i] != nil {
			if err := o.PMMAgent[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listAgentsOk" + "." + "pmm_agent" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListAgentsOKBody) validatePostgresExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.PostgresExporter) { // not required
		return nil
	}

	for i := 0; i < len(o.PostgresExporter); i++ {
		if swag.IsZero(o.PostgresExporter[i]) { // not required
			continue
		}

		if o.PostgresExporter[i] != nil {
			if err := o.PostgresExporter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listAgentsOk" + "." + "postgres_exporter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListAgentsOKBody) validateQANMongoProfilerAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.QANMongoProfilerAgent) { // not required
		return nil
	}

	for i := 0; i < len(o.QANMongoProfilerAgent); i++ {
		if swag.IsZero(o.QANMongoProfilerAgent[i]) { // not required
			continue
		}

		if o.QANMongoProfilerAgent[i] != nil {
			if err := o.QANMongoProfilerAgent[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listAgentsOk" + "." + "qan_mongo_profiler_agent" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListAgentsOKBody) validateQANMysqlPerfschemaAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.QANMysqlPerfschemaAgent) { // not required
		return nil
	}

	for i := 0; i < len(o.QANMysqlPerfschemaAgent); i++ {
		if swag.IsZero(o.QANMysqlPerfschemaAgent[i]) { // not required
			continue
		}

		if o.QANMysqlPerfschemaAgent[i] != nil {
			if err := o.QANMysqlPerfschemaAgent[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listAgentsOk" + "." + "qan_mysql_perfschema_agent" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListAgentsOKBody) validateRDSExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.RDSExporter) { // not required
		return nil
	}

	for i := 0; i < len(o.RDSExporter); i++ {
		if swag.IsZero(o.RDSExporter[i]) { // not required
			continue
		}

		if o.RDSExporter[i] != nil {
			if err := o.RDSExporter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listAgentsOk" + "." + "rds_exporter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListAgentsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAgentsOKBody) UnmarshalBinary(b []byte) error {
	var res ListAgentsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MongodbExporterItems0 MongoDBExporter runs on Generic or Container Node and exposes MongoDB Service metrics.
swagger:model MongodbExporterItems0
*/
type MongodbExporterItems0 struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// MongoDB password for scraping metrics.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// MongoDB username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this mongodb exporter items0
func (o *MongodbExporterItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mongodbExporterItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mongodbExporterItems0TypeStatusPropEnum = append(mongodbExporterItems0TypeStatusPropEnum, v)
	}
}

const (

	// MongodbExporterItems0StatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	MongodbExporterItems0StatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// MongodbExporterItems0StatusSTARTING captures enum value "STARTING"
	MongodbExporterItems0StatusSTARTING string = "STARTING"

	// MongodbExporterItems0StatusRUNNING captures enum value "RUNNING"
	MongodbExporterItems0StatusRUNNING string = "RUNNING"

	// MongodbExporterItems0StatusWAITING captures enum value "WAITING"
	MongodbExporterItems0StatusWAITING string = "WAITING"

	// MongodbExporterItems0StatusSTOPPING captures enum value "STOPPING"
	MongodbExporterItems0StatusSTOPPING string = "STOPPING"

	// MongodbExporterItems0StatusDONE captures enum value "DONE"
	MongodbExporterItems0StatusDONE string = "DONE"
)

// prop value enum
func (o *MongodbExporterItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, mongodbExporterItems0TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *MongodbExporterItems0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *MongodbExporterItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MongodbExporterItems0) UnmarshalBinary(b []byte) error {
	var res MongodbExporterItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MysqldExporterItems0 MySQLdExporter runs on Generic or Container Node and exposes MySQL and AmazonRDSMySQL Service metrics.
swagger:model MysqldExporterItems0
*/
type MysqldExporterItems0 struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// MySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// MySQL username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this mysqld exporter items0
func (o *MysqldExporterItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mysqldExporterItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mysqldExporterItems0TypeStatusPropEnum = append(mysqldExporterItems0TypeStatusPropEnum, v)
	}
}

const (

	// MysqldExporterItems0StatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	MysqldExporterItems0StatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// MysqldExporterItems0StatusSTARTING captures enum value "STARTING"
	MysqldExporterItems0StatusSTARTING string = "STARTING"

	// MysqldExporterItems0StatusRUNNING captures enum value "RUNNING"
	MysqldExporterItems0StatusRUNNING string = "RUNNING"

	// MysqldExporterItems0StatusWAITING captures enum value "WAITING"
	MysqldExporterItems0StatusWAITING string = "WAITING"

	// MysqldExporterItems0StatusSTOPPING captures enum value "STOPPING"
	MysqldExporterItems0StatusSTOPPING string = "STOPPING"

	// MysqldExporterItems0StatusDONE captures enum value "DONE"
	MysqldExporterItems0StatusDONE string = "DONE"
)

// prop value enum
func (o *MysqldExporterItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, mysqldExporterItems0TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *MysqldExporterItems0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *MysqldExporterItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MysqldExporterItems0) UnmarshalBinary(b []byte) error {
	var res MysqldExporterItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*NodeExporterItems0 NodeExporter runs on Generic on Container Node and exposes its metrics.
swagger:model NodeExporterItems0
*/
type NodeExporterItems0 struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`
}

// Validate validates this node exporter items0
func (o *NodeExporterItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var nodeExporterItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeExporterItems0TypeStatusPropEnum = append(nodeExporterItems0TypeStatusPropEnum, v)
	}
}

const (

	// NodeExporterItems0StatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	NodeExporterItems0StatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// NodeExporterItems0StatusSTARTING captures enum value "STARTING"
	NodeExporterItems0StatusSTARTING string = "STARTING"

	// NodeExporterItems0StatusRUNNING captures enum value "RUNNING"
	NodeExporterItems0StatusRUNNING string = "RUNNING"

	// NodeExporterItems0StatusWAITING captures enum value "WAITING"
	NodeExporterItems0StatusWAITING string = "WAITING"

	// NodeExporterItems0StatusSTOPPING captures enum value "STOPPING"
	NodeExporterItems0StatusSTOPPING string = "STOPPING"

	// NodeExporterItems0StatusDONE captures enum value "DONE"
	NodeExporterItems0StatusDONE string = "DONE"
)

// prop value enum
func (o *NodeExporterItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nodeExporterItems0TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *NodeExporterItems0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NodeExporterItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NodeExporterItems0) UnmarshalBinary(b []byte) error {
	var res NodeExporterItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PMMAgentItems0 PMMAgent runs on Generic on Container Node.
swagger:model PMMAgentItems0
*/
type PMMAgentItems0 struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// True if Agent is running and connected to pmm-managed.
	Connected bool `json:"connected,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`
}

// Validate validates this PMM agent items0
func (o *PMMAgentItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PMMAgentItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PMMAgentItems0) UnmarshalBinary(b []byte) error {
	var res PMMAgentItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostgresExporterItems0 PostgresExporter runs on Generic or Container Node and exposes PostgreSQL Service metrics.
swagger:model PostgresExporterItems0
*/
type PostgresExporterItems0 struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// PostgreSQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// PostgreSQL username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this postgres exporter items0
func (o *PostgresExporterItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postgresExporterItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postgresExporterItems0TypeStatusPropEnum = append(postgresExporterItems0TypeStatusPropEnum, v)
	}
}

const (

	// PostgresExporterItems0StatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	PostgresExporterItems0StatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// PostgresExporterItems0StatusSTARTING captures enum value "STARTING"
	PostgresExporterItems0StatusSTARTING string = "STARTING"

	// PostgresExporterItems0StatusRUNNING captures enum value "RUNNING"
	PostgresExporterItems0StatusRUNNING string = "RUNNING"

	// PostgresExporterItems0StatusWAITING captures enum value "WAITING"
	PostgresExporterItems0StatusWAITING string = "WAITING"

	// PostgresExporterItems0StatusSTOPPING captures enum value "STOPPING"
	PostgresExporterItems0StatusSTOPPING string = "STOPPING"

	// PostgresExporterItems0StatusDONE captures enum value "DONE"
	PostgresExporterItems0StatusDONE string = "DONE"
)

// prop value enum
func (o *PostgresExporterItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postgresExporterItems0TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PostgresExporterItems0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostgresExporterItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostgresExporterItems0) UnmarshalBinary(b []byte) error {
	var res PostgresExporterItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*QANMongoProfilerAgentItems0 QANMongoProfilerAgent runs within pmm-agent and sends MongoDB Query Analytics data to the PMM Server.
swagger:model QANMongoProfilerAgentItems0
*/
type QANMongoProfilerAgentItems0 struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// MySQL password for getting performance data.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// MySQL username for getting performance data.
	Username string `json:"username,omitempty"`
}

// Validate validates this QAN mongo profiler agent items0
func (o *QANMongoProfilerAgentItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var qanMongoProfilerAgentItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		qanMongoProfilerAgentItems0TypeStatusPropEnum = append(qanMongoProfilerAgentItems0TypeStatusPropEnum, v)
	}
}

const (

	// QANMongoProfilerAgentItems0StatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	QANMongoProfilerAgentItems0StatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// QANMongoProfilerAgentItems0StatusSTARTING captures enum value "STARTING"
	QANMongoProfilerAgentItems0StatusSTARTING string = "STARTING"

	// QANMongoProfilerAgentItems0StatusRUNNING captures enum value "RUNNING"
	QANMongoProfilerAgentItems0StatusRUNNING string = "RUNNING"

	// QANMongoProfilerAgentItems0StatusWAITING captures enum value "WAITING"
	QANMongoProfilerAgentItems0StatusWAITING string = "WAITING"

	// QANMongoProfilerAgentItems0StatusSTOPPING captures enum value "STOPPING"
	QANMongoProfilerAgentItems0StatusSTOPPING string = "STOPPING"

	// QANMongoProfilerAgentItems0StatusDONE captures enum value "DONE"
	QANMongoProfilerAgentItems0StatusDONE string = "DONE"
)

// prop value enum
func (o *QANMongoProfilerAgentItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, qanMongoProfilerAgentItems0TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *QANMongoProfilerAgentItems0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QANMongoProfilerAgentItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QANMongoProfilerAgentItems0) UnmarshalBinary(b []byte) error {
	var res QANMongoProfilerAgentItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*QANMysqlPerfschemaAgentItems0 QANMySQLPerfSchemaAgent runs within pmm-agent and sends MySQL Query Analytics data to the PMM Server.
swagger:model QANMysqlPerfschemaAgentItems0
*/
type QANMysqlPerfschemaAgentItems0 struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// MySQL password for getting performance data.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// MySQL username for getting performance data.
	Username string `json:"username,omitempty"`
}

// Validate validates this QAN mysql perfschema agent items0
func (o *QANMysqlPerfschemaAgentItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var qanMysqlPerfschemaAgentItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		qanMysqlPerfschemaAgentItems0TypeStatusPropEnum = append(qanMysqlPerfschemaAgentItems0TypeStatusPropEnum, v)
	}
}

const (

	// QANMysqlPerfschemaAgentItems0StatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	QANMysqlPerfschemaAgentItems0StatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// QANMysqlPerfschemaAgentItems0StatusSTARTING captures enum value "STARTING"
	QANMysqlPerfschemaAgentItems0StatusSTARTING string = "STARTING"

	// QANMysqlPerfschemaAgentItems0StatusRUNNING captures enum value "RUNNING"
	QANMysqlPerfschemaAgentItems0StatusRUNNING string = "RUNNING"

	// QANMysqlPerfschemaAgentItems0StatusWAITING captures enum value "WAITING"
	QANMysqlPerfschemaAgentItems0StatusWAITING string = "WAITING"

	// QANMysqlPerfschemaAgentItems0StatusSTOPPING captures enum value "STOPPING"
	QANMysqlPerfschemaAgentItems0StatusSTOPPING string = "STOPPING"

	// QANMysqlPerfschemaAgentItems0StatusDONE captures enum value "DONE"
	QANMysqlPerfschemaAgentItems0StatusDONE string = "DONE"
)

// prop value enum
func (o *QANMysqlPerfschemaAgentItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, qanMysqlPerfschemaAgentItems0TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *QANMysqlPerfschemaAgentItems0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QANMysqlPerfschemaAgentItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QANMysqlPerfschemaAgentItems0) UnmarshalBinary(b []byte) error {
	var res QANMysqlPerfschemaAgentItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RDSExporterItems0 RDSExporter runs on Generic or Container Node and exposes RemoteAmazonRDS Node and AmazonRDSMySQL Service metrics.
swagger:model RDSExporterItems0
*/
type RDSExporterItems0 struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// A list of Service identifiers (Node identifiers are extracted from Services).
	ServiceIds []string `json:"service_ids"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`
}

// Validate validates this RDS exporter items0
func (o *RDSExporterItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rdsExporterItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rdsExporterItems0TypeStatusPropEnum = append(rdsExporterItems0TypeStatusPropEnum, v)
	}
}

const (

	// RDSExporterItems0StatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	RDSExporterItems0StatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// RDSExporterItems0StatusSTARTING captures enum value "STARTING"
	RDSExporterItems0StatusSTARTING string = "STARTING"

	// RDSExporterItems0StatusRUNNING captures enum value "RUNNING"
	RDSExporterItems0StatusRUNNING string = "RUNNING"

	// RDSExporterItems0StatusWAITING captures enum value "WAITING"
	RDSExporterItems0StatusWAITING string = "WAITING"

	// RDSExporterItems0StatusSTOPPING captures enum value "STOPPING"
	RDSExporterItems0StatusSTOPPING string = "STOPPING"

	// RDSExporterItems0StatusDONE captures enum value "DONE"
	RDSExporterItems0StatusDONE string = "DONE"
)

// prop value enum
func (o *RDSExporterItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, rdsExporterItems0TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *RDSExporterItems0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RDSExporterItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RDSExporterItems0) UnmarshalBinary(b []byte) error {
	var res RDSExporterItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
