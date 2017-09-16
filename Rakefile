#
#
#

namespace :golang do
  desc 'Format all Golang files in the current and subdirectories'
  task :fmt do
    sh 'go fmt ./...'
  end

  desc 'Build the current Golang project'
  task :build do
    sh 'go build'
  end

  desc 'Install the current Golang project'
  task :install do
    sh 'go install'
  end

  desc 'Run GoDoc HTTP Server, REQUIRES GODOC TO BE INSTALLED'
  task :doc do
    puts 'Running GoDoc on http://localhost:6060'
    sh 'godoc -http=:6060'
  end
end
