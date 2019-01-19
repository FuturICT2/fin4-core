module UserLogin.ViewLogin exposing (..)

import Html exposing (..)
import Html.Attributes exposing (href)
import Html.Events exposing (on, keyCode)
import Debug exposing (log)
import RemoteData exposing (RemoteData)


--import Html.Events exposing (..)

import Html.Attributes exposing (style)
import UserLogin.Msg exposing (Msg(..))
import UserLogin.Model exposing (Model)
import Material.Textfield as Textfield
import Material.Button as Button
import Material.Options as Options
import Material.Elevation as Elevation
import Common.Error exposing (renderHttpError)
import Main.Routing exposing (loginPath, signupPath, forgotPassPath)
import Json.Decode as JD


onKeyDown : (Int -> msg) -> Options.Property c msg
onKeyDown tagger =
    Options.on "keydown" (JD.map tagger keyCode)


render : Model -> Html Msg
render model =
    div [ loginWrapStyle ]
        [ h1 [] [ text "Log in" ]
        , div []
            [ Textfield.render Mdl
                [ 1 ]
                model.mdl
                [ Textfield.label "Email"
                , Textfield.floatingLabel
                , Textfield.email
                , Textfield.value model.email
                , Options.onInput SetEmail
                ]
                []
            ]
        , div []
            [ Textfield.render Mdl
                [ 2 ]
                model.mdl
                [ Textfield.label "Password"
                , Textfield.floatingLabel
                , Textfield.password
                , Textfield.value model.password
                , Options.onInput SetPassword
                , onKeyDown OnKeyDownLogin
                ]
                []
            ]
        , renderHttpError model.loginError
        , div []
            [ Button.render Mdl
                [ 3 ]
                model.mdl
                [ Button.ripple
                , Button.colored
                , Button.raised
                , Options.onClick DoLogin
                , Options.disabled model.isLoggingIn
                ]
                [ text "Log in" ]
            ]
        , div [ helperStyle ]
            [ text "Not a member yet?"
            , a [ href signupPath ] [ text " Sign up now" ]
            ]
        , div [ helperStyle ]
            [ a [ href forgotPassPath ] [ text "Forgot Password?" ]
            ]
        ]


loginWrapStyle : Attribute Msg
loginWrapStyle =
    style
        [ ( "margin", "0 auto" )
        , ( "max-width", "800px" )
        , ( "text-align", "center" )
        , ( "padding-top", "100px" )
        ]


helperStyle =
    style
        [ ( "padding-top", "30px" )
        ]
