module UserLogin.Update exposing (login, update)

import Debug
import Http
import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Json.Encode as JE
import Main.Context exposing (Context, isUserId)
import Main.User exposing (User, userDecoder)
import Material
import Navigation exposing (newUrl)
import RemoteData exposing (RemoteData)
import UserLogin.Command exposing (loginCmd, signupCmd)
import UserLogin.Model exposing (Model)
import UserLogin.Msg exposing (Msg(..))


login ctx model =
    { model | user = Nothing, isLoggingIn = True, loginError = Nothing } ! [ loginCmd ctx model ]


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        SetName value ->
            { model | name = value } ! []

        DoSignup ->
            { model
                | user = Nothing
                , isSigningUp = True
                , signupError = Nothing
            }
                ! [ signupCmd ctx model ]

        OnSignupResponse res ->
            case res of
                Ok user ->
                    { model
                        | user = Just user
                        , isSigningUp = False
                        , signupError = Nothing
                        , isNewUser = True
                    }
                        ! [ newUrl "#welcome" ]

                Err error ->
                    { model | isSigningUp = False, signupError = Just error } ! []

        OnKeyDownLogin keyCode ->
            if keyCode == 13 then
                login ctx model

            else
                model ! []

        DoLogin ->
            login ctx model

        OnLoginResponse res ->
            case res of
                Ok user ->
                    { model
                        | user = Just user
                        , isLoggingIn = False
                        , loginError = Nothing
                    }
                        ! []

                Err error ->
                    { model | isLoggingIn = False, loginError = Just error } ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
