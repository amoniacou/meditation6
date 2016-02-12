module Meditation
  class Person
    include ::Mongoid::Document
    field :first_name, type: ::String
    field :last_name, type: ::String
    field :email, type: ::String
    field :company, type: ::String
    index({ email: 1 }, { unique: true})
    validates :first_name, :last_name, :email, :company, presence: true
    validates :email, uniqueness: true
  end
end
