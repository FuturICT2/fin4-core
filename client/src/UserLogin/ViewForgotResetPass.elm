module UserLogin.ViewForgotResetPass exposing (..)

import Html exposing (..)
import Html.Attributes exposing (style, attribute, href)
import Material.Options as Options exposing (css)
import Material.Typography as Typo
import Material.Textfield as Textfield
import Material.Button as Button
import Main.Context exposing (Context)
import UserLogin.Model exposing (Model)
import UserLogin.Msg exposing (Msg(..), ForgotResetField(..))
import Common.Error exposing (renderHttpError)
import Common.SuccessMsg as SuccessMsg
import Main.Routing exposing (loginPath)


render : Context -> Model -> Html Msg
render ctx model =
    div [ style [ ( "margin", "100px 0" ), ( "text-align", "center" ) ] ]
        [ Options.styled p
            [ Typo.subhead ]
            [ text "Reset your password" ]
        , case model.forgotReset.success of
            True ->
                div []
                    [ SuccessMsg.render "Your password was changed correctly"
                    , div [] [ text "Go to ", a [ href loginPath ] [ text "log in" ] ]
                    ]

            False ->
                renderForm ctx model
        ]


renderForm : Context -> Model -> Html Msg
renderForm ctx model =
    let
        token =
            case model.forgotReset.token of
                Just token ->
                    token

                Nothing ->
                    "!"

        password =
            model.forgotReset.password

        confirm =
            model.forgotReset.confirm
    in
        div []
            [ div []
                [ Textfield.render Mdl
                    [ 10101 ]
                    model.mdl
                    [ Textfield.label "New password"
                    , Textfield.floatingLabel
                    , Textfield.password
                    , Options.onInput <| SetForgotResetField ForgotResetFieldPassword
                    , Textfield.value password
                    ]
                    []
                ]
            , div []
                [ Textfield.render Mdl
                    [ 10102 ]
                    model.mdl
                    [ Textfield.label "Confirm new password"
                    , Textfield.floatingLabel
                    , Textfield.password
                    , Options.onInput <| SetForgotResetField ForgotResetFieldConfirm
                    , Textfield.value confirm
                    ]
                    []
                ]
            , renderHttpError model.forgotReset.error
            , div []
                [ Button.render Mdl
                    [ 10104 ]
                    model.mdl
                    [ Button.raised
                    , Button.colored
                    , Button.ripple
                    , Options.onClick PostForgotResetPassword
                    , Options.disabled (model.forgotReset.isPosting || (password /= confirm && (String.length password) > 0))
                    ]
                    [ (if model.forgotReset.isPosting then
                        i [ attribute "class" "fa fa-spin fa-spinner" ] []
                       else
                        i [] []
                      )
                    , text
                        "Update password"
                    ]
                ]
            ]
