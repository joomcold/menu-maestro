class Api::ApiController < ApplicationController
  before_action :authenticate_user!

  rescue_from ActiveRecord::RecordNotFound, with: :record_not_found
  rescue_from ActionController::ParameterMissing, with: :parameter_missing

  private

  def record_not_found
    render(json: { error: 'Record not found' }, status: :not_found)
  end

  def parameter_missing(error)
    render(json: { error: "Parameter missing: #{error.param}" }, status: :unprocessable_entity)
  end
end
