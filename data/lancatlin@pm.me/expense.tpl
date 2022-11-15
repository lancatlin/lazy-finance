{{ .Date }} {{ with .Name }}{{ . }}{{ else }}{{ .Destination }}{{ end }}
    {{ .Destination }}      ${{ .Amount }}
    {{ with .Source }}{{ . }}{{ else }}cash{{ end }}

