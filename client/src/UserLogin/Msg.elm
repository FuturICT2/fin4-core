module UserLogin.Msg exposing (Msg(..))

import Http exposing (..)
import Main.User exposing (User)
import Material
import RemoteData exposing (WebData)


type Msg
    = Mdl (Material.Msg Msg)
    | SetName String
    | OnKeyDownLogin Int
    | DoSignup
    | OnSignupResponse (Result Http.Error User)
    | DoLogin
    | OnLoginResponse (Result Http.Error User)
