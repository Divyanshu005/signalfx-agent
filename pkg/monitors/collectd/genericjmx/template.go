// +build linux

package genericjmx

// AUTOGENERATED BY scripts/collectd-template-to-go.  DO NOT EDIT!!

import (
	"text/template"

	"github.com/signalfx/signalfx-agent/pkg/monitors/collectd"
)

// CollectdTemplate is a template for a genericjmx collectd config file
var CollectdTemplate = template.Must(collectd.InjectTemplateFuncs(template.New("genericjmx")).Parse(`
<Plugin java>
  <Plugin "GenericJMX">

    {{range $name, $_ := .MBeanDefinitions}}
    <MBean "{{$name}}">
      ObjectName "{{.ObjectName}}"
      {{with .InstancePrefix -}}
      InstancePrefix "{{.}}"
      {{end -}}
      {{range .InstanceFrom -}}
      InstanceFrom "{{.}}"
      {{end -}}
      {{range .Dimensions -}}
      Dimension "{{.}}"
      {{end -}}
      {{range .Values}}
      <Value>
        {{with .InstancePrefix -}}
        InstancePrefix "{{.}}"
        {{- end}}
        {{range .InstanceFrom -}}
        InstanceFrom "{{.}}"
        {{end -}}
        Type "{{.Type}}"
        Table {{withDefault .Table "false"}}
        {{range .Attributes}}
        Attribute "{{.}}"
        {{else}}
        Attribute "{{.Attribute}}"
        {{end}}
      </Value>
      {{end}}
    </MBean>
    {{end}}

    <Connection>
      {{with .ServiceURL}}
      ServiceURL "{{renderValue . $}}"
      {{- end}}
      {{with .InstancePrefix -}}
      InstancePrefix "{{.}}"
      {{- end}}
      ServiceName "{{.ServiceName}}"
      {{with .Username -}}
      User "{{.}}"
      {{- end}}
      {{- with .Password}}
      Password "{{.}}"
      {{- end}}
      {{range $key, $val := .CustomDimensions}}
      CustomDimension "{{$key}}" "{{$val}}"
      {{end}}
      CustomDimension "monitorID" "{{.MonitorID}}"
      {{range .MBeansToCollect}}
      Collect "{{.}}"
      {{- end}}
    </Connection>
  </Plugin>
</Plugin>
`)).Option("missingkey=error")