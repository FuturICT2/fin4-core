module UserLogin.Model exposing (Model, init)

import Http
import Json.Decode as JD
import Main.User exposing (User)
import Material
import RemoteData exposing (RemoteData, WebData)


type alias Model =
    { mdl : Material.Model
    , name : String
    , user : Maybe User
    , isSigningUp : Bool
    , signupError : Maybe Http.Error
    , isLoggingIn : Bool
    , loginError : Maybe Http.Error
    , isNewUser : Bool
    }


init : Model
init =
    { mdl = Material.model
    , name = ""
    , user = Nothing
    , isSigningUp = False
    , signupError = Nothing
    , isLoggingIn = False
    , loginError = Nothing
    , isNewUser = False
    }
