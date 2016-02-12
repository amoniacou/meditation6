module Meditation
  class Server < ::Sinatra::Base
    get '/persons' do
      json ::Meditation::Person.all.to_a
    end
    post '/persons' do
      person = ::Meditation::Person.new(JSON.parse(request.body.read))
      if person.save
        status 201
        json person
      else
        status 400
        json errors: person.errors
      end
    end
  end
end

