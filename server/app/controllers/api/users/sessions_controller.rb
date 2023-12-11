class Api::Users::SessionsController < Devise::SessionsController
  respond_to :json

  private

  # Login
  def respond_with(_resource, _opts = {})
    render(json: { message: 'Logged in successfully.', user: current_user }, status: :ok)
  end

  # Logout
  def respond_to_on_destroy
    raise(StandardError, 'Authorization cannot be blank') if request.headers['Authorization'].blank?

    log_out_success && return if user_from_token

    log_out_failed
  end

  def log_out_success
    render(json: { message: 'Logged out successfully.' }, status: :ok)
  end

  def log_out_failed
    render(json: { message: 'Something went wrong.' }, status: :unauthorized)
  end

  def user_from_token
    jwt_payload = JWT.decode(request.headers['Authorization'].split[1],
                             Rails.application.credentials.devise[:jwt_secret_key]).first
    current_user = User.find(jwt_payload['sub'].to_s)

    raise(JWT::VerificationError, 'The user has no session') if current_user&.jti != jwt_payload['jti']

    current_user
  end
end
