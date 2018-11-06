module CreateToken.StepName exposing (render)

import Common.Error exposing (renderHttpError)
import Common.Spinner as Spinner
import CreateToken.Msg exposing (Msg(..))
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (keyCode)
import Json.Decode as JD
import Material
import Material.Card as Card
import Material.Chip as Chip
import Material.Options as Options exposing (css)
import Material.Textfield as Textfield


render ctx model =
    div
        []
        [ div [ margin "14" "0" ]
            [ Textfield.render Mdl
                [ 2 ]
                model.mdl
                [ Textfield.label "Name"
                , Textfield.floatingLabel
                , Textfield.text_
                , Options.onInput SetName
                , Textfield.value model.name
                ]
                []
            ]

        -- , div [ style [ ( "15px", "0" ) ] ]
        --     [ Textfield.render Mdl
        --         [ 1 ]
        --         model.mdl
        --         [ css "width" "200px"
        --         , Textfield.label "Total supply"
        --         , Textfield.floatingLabel
        --         , Textfield.text_
        --         , Textfield.value model.shares
        --         , Options.onInput SetShares
        --         ]
        --         []
        --     ]
        , div
            []
            [ Textfield.render Mdl
                [ 3 ]
                model.mdl
                [ Textfield.label "Symbol"
                , Textfield.floatingLabel
                , Textfield.text_
                , Options.onInput SetSymbol
                , Textfield.value model.symbol
                ]
                []
            ]
        , div []
            [ Textfield.render
                Mdl
                [ 4 ]
                model.mdl
                [ Textfield.label "Purpose"
                , Textfield.floatingLabel
                , Textfield.textarea
                , Textfield.rows 3
                , Options.onInput SetDescription
                , Textfield.value model.description
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
