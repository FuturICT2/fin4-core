module CreateToken.StepName exposing (render)

import Common.Error exposing (renderHttpError)
import Common.Spinner as Spinner
import CreateToken.Msg exposing (Msg(..))
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onInput)
import Json.Decode as JD
import Material
import Material.Card as Card
import Material.Chip as Chip
import Material.Options as Options exposing (css)
import Material.Textfield as Textfield


render ctx model =
    div
        []
        [ div [ style [ ( "margin", "30px 0 0 0" ) ] ]
            [ p
                [ style
                    [ ( "margin", "0" )
                    , ( "padding-left", "2px" )
                    ]
                ]
                [ text "Enter name" ]
            , input
                [ inputStyle
                , placeholder "e.g Moon"
                , value model.name
                , onInput SetName
                ]
                []
            ]
        , div [ style [ ( "margin", "30px 0 0 0" ) ] ]
            [ p
                [ style
                    [ ( "margin", "0" )
                    , ( "padding-left", "2px" )
                    ]
                ]
                [ text "Enter symbol" ]
            , input
                [ inputStyle
                , placeholder "e.g MNN"
                , value model.symbol
                , onInput SetSymbol
                ]
                []
            ]
        , div
            [ style
                [ ( "margin", "15px 0 0 0" )
                ]
            ]
            [ p
                [ style
                    [ ( "margin", " 0" )
                    , ( "padding-left", "2px" )
                    ]
                ]
                [ text "Enter purpose" ]
            , textarea
                [ textareaStyle
                , value model.description
                , placeholder "e.g Let us go to the moon"
                , onInput SetDescription
                , rows 5
                ]
                []
            ]
        , renderHttpError model.createTokenError
        , case model.isCreatingToken of
            True ->
                div
                    [ style
                        [ ( "marging-top", "50px" )
                        , ( "text-align", "center" )
                        ]
                    ]
                    [ Spinner.render "Deploying your token to the Ethereum network. This might take few moments."
                    ]

            False ->
                div [] []
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
