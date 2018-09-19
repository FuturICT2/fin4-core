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


view : Model -> Html Msg
view model =
    case model.context.sessionDidLoad of
        False ->
            Spinner.render "..."

        True ->
            div [ containerStyle model.context.window.isMobile ]
                [ ViewHeader.render model
                , ViewMobileNav.render model
                , div [ bodyStyle ] [ ViewBody.render model ]
                , ViewFooter.render model
                ]


containerStyle : Bool -> Attribute a
containerStyle isMobile =
    style <|
        if isMobile then
            [ ( "padding-bottom", "100px" ) ]

        else
            [ ( "padding-top", "40px" )
            , ( "min-width", "950px" )
            , ( "background", "#fcfcfc" )
            ]


bodyStyle : Attribute a
bodyStyle =
    style
        [ ( "min-height", "700px" )
        , ( "padding-left", "10px" )
        , ( "padding-right", "15px" )
        ]
