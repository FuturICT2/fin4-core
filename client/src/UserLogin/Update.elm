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
import UserLogin.Command exposing (loginCmd, requestNewPasswordCmd, requestNewPasswordPostResetCmd, signupCmd)
import UserLogin.Model exposing (Model)
import UserLogin.Msg exposing (ForgotResetField(..), Msg(..))


login ctx model =
    { model | user = Nothing, isLoggingIn = True, loginError = Nothing } ! [ loginCmd ctx model ]


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        SetEmail value ->
            { model | email = value } ! []

        SetName value ->
            { model | name = value } ! []

        SetPassword value ->
            { model | password = value } ! []

        ToggleAgreeToTerms ->
            { model | agreeToTerms = not model.agreeToTerms } ! []

        ShowTermsOfService ->
            { model | showTerms = not model.showTerms } ! []

        ShowPrivacy ->
            { model | showPrivacy = not model.showPrivacy } ! []

        DoSignup ->
            { model
                | user = Nothing
                , isSigningUp = True
                , signupError = Nothing
            }
                ! [ signupCmd ctx False model ]

        DoFastSignup ->
            { model
                | user = Nothing
                , isSigningUp = True
                , signupError = Nothing
            }
                ! [ signupCmd ctx True model ]

        OnSignupResponse res ->
            case res of
                Ok user ->
                    { model
                        | user = Just user
                        , isSigningUp = False
                        , signupError = Nothing
                        , isNewUser = True
                    }
                        ! [ newUrl "#tokens" ]

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

        RequestNewPassword ->
            { model
                | requestPassError = Nothing
                , isRequestingPass = True
            }
                ! [ requestNewPasswordCmd ctx model ]

        OnRequestNewPasswordResponse resp ->
            case resp of
                Ok response ->
                    { model
                        | requestPassResponse = Just response
                        , isRequestingPass = False
                    }
                        ! []

                Err error ->
                    { model
                        | requestPassError = Just error
                        , isRequestingPass = False
                    }
                        ! []

        InitForgotResetPass userId token ->
            let
                forgotReset =
                    model.forgotReset
            in
            { model
                | forgotReset =
                    { forgotReset
                        | userId = Just userId
                        , token = Just token
                    }
            }
                ! []

        SetForgotResetField field value ->
            case field of
                ForgotResetFieldPassword ->
                    let
                        forgotReset =
                            model.forgotReset
                    in
                    { model
                        | forgotReset =
                            { forgotReset
                                | password = value
                            }
                    }
                        ! []

                ForgotResetFieldConfirm ->
                    let
                        forgotReset =
                            model.forgotReset
                    in
                    { model
                        | forgotReset =
                            { forgotReset
                                | confirm = value
                            }
                    }
                        ! []

        PostForgotResetPassword ->
            let
                forgotReset =
                    model.forgotReset
            in
            { model
                | forgotReset =
                    { forgotReset
                        | isPosting = True
                        , error = Nothing
                    }
            }
                ! [ requestNewPasswordPostResetCmd ctx model ]

        OnForgotPassResetResponse resp ->
            let
                forgotReset =
                    model.forgotReset
            in
            case resp of
                Ok response ->
                    { model
                        | forgotReset =
                            { forgotReset
                                | isPosting = False
                                , error = Nothing
                                , success = True
                            }
                    }
                        ! []

                Err err ->
                    { model
                        | forgotReset =
                            { forgotReset
                                | isPosting = False
                                , error = Just err
                            }
                    }
                        ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
