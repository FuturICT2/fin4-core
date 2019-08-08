module UserLogin.Command exposing (encodeLogin, encodeRequestNewPassword, loginCmd, requestNewPasswordCmd, requestNewPasswordPostResetCmd, requestNewPasswordResponseDecoder, signupCmd, withDefault)

import Common.Api exposing (postWithCsrf)
import Http
import HttpBuilder exposing (..)
import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.User exposing (User, userDecoder)
import RemoteData exposing (RemoteData)
import Task exposing (Task)
import UserLogin.Model exposing (Model, RequestNewPasswordResponse)
import UserLogin.Msg exposing (Msg(..))


signupCmd : Context -> Bool -> Model -> Cmd Msg
signupCmd ctx isFastSignup model =
    postWithCsrf ctx
        OnSignupResponse
        "/register"
        (JE.object
            [ ( "email", JE.string model.email )
            , ( "name", JE.string model.name )
            , ( "password", JE.string model.password )
            , ( "agreeToTerms", JE.bool True ) --model.agreeToTerms )
            , ( "isFastSignup", JE.bool isFastSignup )
            ]
        )
        userDecoder


loginCmd : Context -> Model -> Cmd Msg
loginCmd ctx model =
    postWithCsrf ctx
        OnLoginResponse
        "/login"
        (encodeLogin model)
        userDecoder


encodeLogin : Model -> JE.Value
encodeLogin model =
    JE.object
        [ ( "name", JE.string model.email )
        ]


requestNewPasswordCmd : Context -> Model -> Cmd Msg
requestNewPasswordCmd ctx model =
    postWithCsrf ctx
        OnRequestNewPasswordResponse
        "/forgotpass-requests/new"
        (encodeRequestNewPassword model)
        requestNewPasswordResponseDecoder


requestNewPasswordPostResetCmd : Context -> Model -> Cmd Msg
requestNewPasswordPostResetCmd ctx model =
    postWithCsrf ctx
        OnForgotPassResetResponse
        "/forgotpass-requests/reset"
        (JE.object
            [ ( "userId", JE.int <| withDefault 0 model.forgotReset.userId )
            , ( "token", JE.string <| withDefault "" model.forgotReset.token )
            , ( "password", JE.string model.forgotReset.password )
            ]
        )
        (JP.decode {})


withDefault : a -> Maybe a -> a
withDefault default maybeVal =
    case maybeVal of
        Just v ->
            v

        Nothing ->
            default


encodeRequestNewPassword : Model -> JE.Value
encodeRequestNewPassword model =
    JE.object [ ( "email", JE.string model.email ) ]


requestNewPasswordResponseDecoder : JD.Decoder RequestNewPasswordResponse
requestNewPasswordResponseDecoder =
    JP.decode RequestNewPasswordResponse
