require 'faker'
require 'json'

def make_post
  json_headers = {"Content-Type" => "application/json",
                "Accept" => "application/json"}
  params = {
    'first_name' => Faker::Name.first_name, 'last_name' => Faker::Name.last_name,
    'email' => Faker::Internet.email, 'company' => Faker::Company.name
  }
  uri = URI.parse('http://localhost:8080/persons')
  http = Net::HTTP.new(uri.host, uri.port)
  http.post(uri.path, params.to_json, json_headers)
end

amount = (ENV["AMOUNT"] || 1000).to_i
time = ::Time.now
if ENV["THREADS"].to_i > 0
  count = amount / ENV["THREADS"].to_i
  (1..ENV["THREADS"].to_i).map do |i|
    Thread.new(i) { count.times { make_post } }
  end.each {|t| t.join}
else
  amount.times do
    make_post
  end
end
puts ::Time.now - time