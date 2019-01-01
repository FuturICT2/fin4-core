module CreateToken.StepName exposing (render)

import Common.Error exposing (renderHttpError)
import Common.Spinner as Spinner
import CreateToken.Msg exposing (Msg(..))
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onInput)
import Identicon exposing (identicon)
import Json.Decode as JD
import Material
import Material.Card as Card
import Material.Chip as Chip
import Material.Options as Options exposing (css)
import Material.Textfield as Textfield


render ctx model =
    div
        []
        [ div [ identiconStyle ] [ identicon "60px" model.name ]
        , div [ style [ ( "margin", "30px 0 0 0" ) ] ]
            [ p
                [ style
                    [ ( "margin", "0" )
                    , ( "padding-left", "2px" )
                    ]
                ]
                [ text "Enter name" ]
            , input
                [ inputStyle
                , placeholder "e.g SchoolCoin"
                , value model.name
                , onInput SetName
                ]
                []
            ]
        , div []
            [ p
                [ style
                    [ ( "margin", "0" )
                    , ( "padding-left", "2px" )
                    ]
                ]
                [ text "Enter symbol" ]
            , input
                [ inputStyle
                , placeholder "e.g SCL"
                , value model.symbol
                , onInput SetSymbol
                ]
                []
            ]
        , div
            []
            [ p
                [ style
                    [ ( "margin", " 0" )
                    , ( "padding-left", "2px" )
                    ]
                ]
                [ text "Enter action purpose" ]
            , textarea
                [ textareaStyle
                , value model.description
                , placeholder "e.g Take a picture of a school, post it here with it's location and get rewarded with SchoolCoin"
                , onInput SetDescription
                , rows 5
                ]
                []
            ]
        , renderHttpError model.createTokenError
        ]


margin upper sides =
    style [ ( "margin", upper ++ " " ++ sides ) ]


btnStyle : Attribute a
btnStyle =
    style
        [ ( "height", "50px" )
        , ( "background-color", "#f2f2f2" )
        , ( "background-color", "#dedede" )
        , ( "border", "none" )
        , ( "border-top", "1px solid black" )
        , ( "font-family", "Roboto" )
        , ( "font-size", "20px" )
        , ( "font-style", "condensed" )
        , ( "font-weight", "bold" )
        , ( "text-align", "center" )
        , ( "color", "black" )
        , ( "width", "100%" )
        , ( "position", "fixed" )
        , ( "bottom", "0" )
        ]


inputStyle : Attribute a
inputStyle =
    style
        [ ( "border-radius", "8px" )
        , ( "background-color", "#f2f2f2" )
        , ( "height", "50px" )
        , ( "border", "none" )
        , ( "padding", "15px" )
        , ( "width", "100%" )
        , ( "box-sizing", "border-box" )
        , ( "margin-bottom", "15px" )
        ]


textareaStyle : Attribute a
textareaStyle =
    style
        [ ( "border-radius", "8px" )
        , ( "background-color", "#f2f2f2" )
        , ( "border", "none" )
        , ( "padding", "15px" )
        , ( "width", "100%" )
        , ( "box-sizing", "border-box" )
        , ( "margin-bottom", "15px" )
        ]


identiconStyle : Attribute a
identiconStyle =
    style
        [ ( "width", "100px" )
        , ( "height", "100px" )
        , ( "border", "1px solid #ddd" )
        , ( "border-radius", "50%" )
        , ( "margin", "auto" )
        , ( "padding", "20px" )
        ]
