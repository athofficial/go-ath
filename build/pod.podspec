Pod::Spec.new do |spec|
  spec.name         = 'gath'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/athofficial/go-ath'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS ATH Client'
  spec.source       = { :git => 'https://github.com/athofficial/go-ath.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/gath.framework'

	spec.prepare_command = <<-CMD
    curl https://gathstore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/gath.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
