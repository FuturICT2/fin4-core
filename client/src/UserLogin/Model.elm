module UserLogin.Model exposing (Model, RequestNewPasswordResponse, init)

import Http
import Json.Decode as JD
import Main.User exposing (User)
import Material
import RemoteData exposing (RemoteData, WebData)


type alias Model =
    { mdl : Material.Model
    , email : String
    , password : String
    , name : String
    , agreeToTerms : Bool
    , user : Maybe User
    , showTerms : Bool
    , showPrivacy : Bool
    , isSigningUp : Bool
    , signupError : Maybe Http.Error
    , isLoggingIn : Bool
    , loginError : Maybe Http.Error
    , requestPassResponse : Maybe RequestNewPasswordResponse
    , requestPassError : Maybe Http.Error
    , isRequestingPass : Bool
    , isNewUser : Bool
    , forgotReset :
        { token : Maybe String
        , userId : Maybe Int
        , password : String
        , confirm : String
        , isPosting : Bool
        , error : Maybe Http.Error
        , success : Bool
        }
    }


init : Model
init =
    { mdl = Material.model
    , email = ""
    , name = ""
    , password = ""
    , agreeToTerms = False
    , user = Nothing
    , showTerms = False
    , showPrivacy = False
    , isSigningUp = False
    , signupError = Nothing
    , isLoggingIn = False
    , loginError = Nothing
    , requestPassResponse = Nothing
    , requestPassError = Nothing
    , isRequestingPass = False
    , isNewUser = False
    , forgotReset =
        { token = Nothing
        , userId = Nothing
        , password = ""
        , confirm = ""
        , isPosting = False
        , error = Nothing
        , success = False
        }
    }


type alias RequestNewPasswordResponse =
    {}
