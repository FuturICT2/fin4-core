module UserLogin.ViewFastSignup exposing (..)

import Common.Error exposing (renderHttpError)
import Common.Modal as Modal
import Html exposing (..)
import Html.Attributes exposing (href, style)
import Html.Events exposing (onClick)
import Main.Routing exposing (loginPath)
import Material.Button as Button
import Material.Options as Options
import Material.Textfield as Textfield
import Material.Toggles as Toggles
import RemoteData exposing (RemoteData)
import Site.Privacy
import Site.TermsOfService
import UserLogin.Model exposing (Model)
import UserLogin.Msg exposing (Msg(..))


render : Model -> Html Msg
render model =
    div [ wrapStyle ]
        [ h2 [] [ text "Congratulations, you did it!" ]
        , div [style[("padding", "10px"),("max-width", "500px"),("margin", "0 auto")]] [
         text "If you register, you will be rewarded 1000 ANYS. "
        ]
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
        , renderHttpError model.signupError
        , div [] []
        , div []
            [ Button.render Mdl
                [ 10 ]
                model.mdl
                [ Button.ripple
                , Button.colored
                , Button.raised
                , Options.onClick DoFastSignup
                , Options.disabled model.isSigningUp
                ]
                [ text "Sign up" ]
            ], div [ helperStyle ]
                [ text "Already a member? "
                , a [ href loginPath ] [ text "Log in" ]
                ]



        , renderTermsOfService model
        , renderPrivacy model
        ]


renderTermsOfService : Model -> Html Msg
renderTermsOfService model =
    case model.showTerms of
        True ->
            Modal.render ShowTermsOfService "Terms of service" Site.TermsOfService.render

        False ->
            span [] []


renderPrivacy : Model -> Html Msg
renderPrivacy model =
    case model.showPrivacy of
        True ->
            Modal.render ShowPrivacy "Privacy policy" Site.Privacy.render

        False ->
            span [] []


wrapStyle : Attribute Msg
wrapStyle =
    style
        [ ( "margin", "0 auto" )
        , ( "max-width", "800px" )
        , ( "text-align", "center" )
        , ( "padding-top", "100px" )
        ]


checkboxWrapStyle : Attribute Msg
checkboxWrapStyle =
    style
        [ ( "width", "300px" )
        , ( "margin", "0 auto" )
        , ( "height", "60px" )
        ]


helperStyle : Attribute a
helperStyle =
    style
        [ ( "padding-top", "30px" )
        ]
