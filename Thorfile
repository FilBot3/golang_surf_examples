# Thor Homepage: http://whatisthor.com/
# Thor GitHub Wiki: https://github.com/erikhuda/thor/wiki
# Thor RDoc: www.rubydoc.info/github/wycats/thor/Thor/Actions

require 'thor'

class Golang < Thor
  include Thor::Actions

  desc 'fmt', 'Formats all Golang files in the current and subdirectories'
  def fmt
    run('go fmt ./...', {:verbose => true})
  end

  desc 'build', 'Build the current Project'
  def build
    run('go build')
  end

  desc 'install', 'Install the Project to local Golang libraries'
  def install
    run('go install')
  end

  desc 'doc', 'Run the godoc server, REQUIRES GODOC'
  def doc
    puts 'Running GODOC HTTP Server at http://localhost:6060'
    puts 'Requires CTRL+c to Exit, will cause Thor Interrupt.'
    run('godoc -http=:6060')
  end

  default_task :fmt
end
