module UserLogin.ViewForgot exposing (render)

import Html exposing (..)
import Html.Attributes exposing (href)
import Html.Attributes exposing (style)
import RemoteData exposing (RemoteData)
import Material.Textfield as Textfield
import Material.Button as Button
import Material.Options as Options
import Main.Routing exposing (loginPath)
import UserLogin.Msg exposing (Msg(..))
import UserLogin.Model exposing (Model)
import Common.Error exposing (renderHttpError)
import Common.SuccessMsg as SuccessMsg


render : Model -> Html Msg
render model =
    div [ wrapStyle ]
        [ h1 [] [ text "Request a new password" ]
        , case model.requestPassResponse of
            Just response ->
                renderSuccess model

            Nothing ->
                renderForm model
        ]


renderSuccess : Model -> Html Msg
renderSuccess model =
    div []
        [ SuccessMsg.render "We have sent instructions for resetting your password to your email."
        , div [] [ text "Go to ", a [ href loginPath ] [ text "log in" ] ]
        ]


renderForm : Model -> Html Msg
renderForm model =
    div []
        [ div []
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
        , renderHttpError model.requestPassError
        , div []
            [ Button.render Mdl
                [ 3 ]
                model.mdl
                [ Button.ripple
                , Button.colored
                , Button.raised
                , Options.onClick RequestNewPassword
                , Options.disabled model.isRequestingPass
                ]
                [ text "Request a new password" ]
            ]
        , div [ helperStyle ]
            [ text "Go to "
            , a [ href loginPath ] [ text "Log in" ]
            ]
        ]


wrapStyle : Attribute a
wrapStyle =
    style
        [ ( "margin", "0 auto" )
        , ( "max-width", "800px" )
        , ( "text-align", "center" )
        , ( "padding-top", "100px" )
        ]


helperStyle : Attribute a
helperStyle =
    style
        [ ( "padding-top", "30px" )
        ]
