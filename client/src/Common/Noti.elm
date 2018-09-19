module Common.Noti exposing (..)

import Html exposing (..)
import Html.Attributes exposing (style)


render : Bool -> String -> Html msg
render isVisible msg =
    case isVisible of
        False ->
            div [] []

        True ->
            div
                [ style
                    [ ( "position", "fixed" )
                    , ( "top", "60px" )
                    , ( "right", "20px" )
                    , ( "z-index", "100000" )
                    , ( "padding", "5px" )
                    , ( "background", "#484848" )
                    , ( "color", "#efefef" )
                    , ( "border-radius", "5px" )
                    ]
                ]
                [ text msg ]
