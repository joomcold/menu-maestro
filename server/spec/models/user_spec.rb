require 'rails_helper'

RSpec.describe(User) do
  describe 'validations' do
    it { is_expected.to(validate_presence_of(:email)) }
    it { is_expected.to(validate_presence_of(:password)) }
    it { is_expected.to(validate_length_of(:password).is_at_least(6)) }
  end

  # describe 'associations' do
  #   it { is_expected.to(have_many(:orders)) }
  # end
end
