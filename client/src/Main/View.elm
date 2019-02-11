module Main.View exposing (bodyStyle, containerStyle, view)

import Common.Spinner as Spinner
import Html exposing (..)
import Html.Attributes exposing (attribute, href, style)
import Main.Context exposing (isLoggedIn)
import Main.Model exposing (Model)
import Main.Msg exposing (Msg)
import Main.ViewBody as ViewBody
import Main.ViewNavigation as ViewNavigation
import Material.Card as Card
import Material.Options exposing (css)


view : Model -> Html Msg
view model =
    div [ mainStyle model.context.window.isMobile ]
        [ case model.context.sessionDidLoad of
            False ->
                renderBeforeSessionLoaded

            True ->
                renderAfterSessionLoaded model
        ]


renderBeforeSessionLoaded : Html Msg
renderBeforeSessionLoaded =
    div [ style [ ( "text-align", "center" ) ] ]
        [ h3
            [ style
                [ ( "margin-top", "60px" )
                ]
            ]
            [ text "Demo" ]
        , Card.view
            [ css "width" "95px"
            , css "height" "95px"
            , css "background" "url('assets/images/coin1.png') center / cover"
            , css "margin" "45px auto"
            ]
            []
        , Spinner.render "One moment..."
        ]


renderAfterSessionLoaded : Model -> Html Msg
renderAfterSessionLoaded model =
    div
        [ containerStyle model.context.window.isMobile
        ]
        [ case isLoggedIn model.context of
            True ->
                ViewNavigation.render model

            False ->
                text ""
        , div [ bodyStyle model.context.window.isMobile ] [ ViewBody.render model ]
        ]


mainStyle : Bool -> Attribute a
mainStyle isMobile =
    case isMobile of
        True ->
            style <|
                []

        False ->
            style <|
                [ ( "padding-bottom", "50px" )
                , ( "margin", "50px auto" )
                , ( "width", "600px" )
                ]


containerStyle : Bool -> Attribute a
containerStyle isMobile =
    case isMobile of
        True ->
            style <|
                [ ( "padding-bottom", "50px" )
                , ( "margin", "0 0" )
                ]

        False ->
            style <|
                [ ( "padding-bottom", "50px" )
                , ( "margin", "0 auto" )
                , ( "max-width", "500px" )
                ]


bodyStyle : Bool -> Attribute a
bodyStyle isMobile =
    let
        w =
            case isMobile of
                True ->
                    "100%"

                False ->
                    "600px"
    in
    style
        [ ( "width", w )
        , ( "margin", "auto" )
        ]
