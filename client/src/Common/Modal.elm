module Common.Modal exposing (..)

import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Material.Icon as Icon
import Material.Options as Options


render onCloseMsg title content =
    div []
        [ div [ maskStyle ]
            [ div [ modalStyle ]
                [ div [ closeStyle ]
                    [ Icon.view "clear"
                        [ Icon.size24
                        , Options.onClick onCloseMsg
                        ]
                    ]
                , div [ modalBodyStyle ]
                    [ header title
                    , div [ contentStyle (title /= "") ] [ content ]
                    ]
                ]
            ]
        ]


header title =
    case title == "" of
        True ->
            div [] []

        False ->
            div [ headerStyle ] [ text title ]


maskStyle : Attribute a
maskStyle =
    style
        [ ( "background-color", "rgba(0,0,0,0.8)" )
        , ( "position", "fixed" )
        , ( "top", "0" )
        , ( "left", "0" )
        , ( "width", "100%" )
        , ( "height", "100%" )
        , ( "z-index", "10001" )
        ]


modalStyle : Attribute a
modalStyle =
    style
        [ ( "background-color", "rgba(255,255,255,1.0)" )
        , ( "position", "fixed" )
        , ( "top", "50%" )
        , ( "left", "50%" )
        , ( "transform", "translate(-50%, -50%)" )
        , ( "width", "700px" )
        , ( "max-width", "100%" )
        , ( "height", "500px" )
        , ( "max-height", "100%" )
        , ( "box-sizing", "border-box" )
        , ( "padding", "10px" )
        , ( "border-radius", "5px" )
        , ( "box-shadow", "1px 1px 5px rgba(0,0,0,0.5)" )
        , ( "z-index", "11" )
        ]


modalBodyStyle =
    style
        []


headerStyle =
    style
        [ ( "font-size", "20px" )
        , ( "font-weight", "bold" )
        , ( "text-align", "center" )
        , ( "padding", "15px" )
        ]


contentStyle : Bool -> Attribute a
contentStyle hasTitle =
    style
        [ ( "overflow", "auto" )
        , ( "position", "absolute" )
        , ( "top"
          , if hasTitle then
                "60px"
            else
                "40px"
          )
        , ( "bottom", "0" )
        , ( "left", "0" )
        , ( "width", "100%" )
        , ( "padding", "0px 20px" )
        , ( "text-align", "justify" )
        ]


closeStyle : Attribute a
closeStyle =
    style [ ( "float", "right" ), ( "cursor", "pointer" ) ]
