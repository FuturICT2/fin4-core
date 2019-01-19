module UserLogin.Msg exposing (ForgotResetField(..), Msg(..))

import Http exposing (..)
import Main.User exposing (User)
import Material
import RemoteData exposing (WebData)
import UserLogin.Model exposing (RequestNewPasswordResponse)


type ForgotResetField
    = ForgotResetFieldPassword
    | ForgotResetFieldConfirm


type Msg
    = Mdl (Material.Msg Msg)
    | SetEmail String
    | SetName String
    | SetPassword String
    | OnKeyDownLogin Int
    | ToggleAgreeToTerms
    | ShowTermsOfService
    | ShowPrivacy
    | DoSignup
    | DoFastSignup
    | OnSignupResponse (Result Http.Error User)
    | DoLogin
    | OnLoginResponse (Result Http.Error User)
    | RequestNewPassword
    | OnRequestNewPasswordResponse (Result Http.Error RequestNewPasswordResponse)
    | InitForgotResetPass Int String
    | SetForgotResetField ForgotResetField String
    | PostForgotResetPassword
    | OnForgotPassResetResponse (Result Http.Error RequestNewPasswordResponse)
