module Main.View exposing (bodyStyle, containerStyle, view)

import Common.Spinner as Spinner
import Html exposing (..)
import Html.Attributes exposing (attribute, href, style)
import Main.Model exposing (Model)
import Main.Msg exposing (Msg)
import Main.ViewBody as ViewBody
import Main.ViewFooter as ViewFooter
import Main.ViewHeader as ViewHeader
import Main.ViewMobileNav as ViewMobileNav
import Material.Card as Card
import Material.Options exposing (css)


view : Model -> Html Msg
view model =
    case model.context.sessionDidLoad of
        False ->
            div [ style [ ( "text-align", "center" ) ] ]
                [ Card.view
                    [ css "width" "95px"
                    , css "height" "95px"
                    , css "background" "url('https://trello-attachments.s3.amazonaws.com/5b39f1d06f761ae7c1c7d22c/5b9f76d5afa89e794357c6d9/6b63377efe7ee3d2e6d1b2af22ca51b7/coin1.png') center / cover"
                    , css "margin" "15px auto"
                    ]
                    []
                , Spinner.render "One moment..."
                ]

        True ->
            div
                [ containerStyle model.context.window.isMobile ]
                [ ViewHeader.render model
                , ViewMobileNav.render model
                , div [ bodyStyle ] [ ViewBody.render model ]
                , ViewFooter.render model
                ]


containerStyle : Bool -> Attribute a
containerStyle isMobile =
    style <| [ ( "padding-bottom", "100px" ) ]



-- if isMobile then
--     [ ( "padding-bottom", "100px" ) ]
--
-- else
--     [ ( "padding-top", "40px" )
--     , ( "min-width", "950px" )
--     , ( "background", "#fcfcfc" )
--     ]


bodyStyle : Attribute a
bodyStyle =
    style
        [ ( "min-height", "700px" )
        , ( "padding-left", "10px" )
        , ( "padding-right", "15px" )
        ]
