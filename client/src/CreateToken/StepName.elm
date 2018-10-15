module CreateToken.StepName exposing (render)

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


categories =
    [ "animals", "nature", "tech", "climate", "health" ]


render ctx model =
    div
        []
        [ Card.view
            [ css "width" "95px"
            , css "height" "95px"
            , css "background" "url('http://www.slate.com/content/dam/slate/articles/health_and_science/science/2017/06/170621_SCI_TreePlantingHoax.jpg.CROP.promo-xlarge2.jpg') center / cover"
            , css "margin" "15px auto"
            ]
            []
        , div [ margin "30px" "0" ]
            [ Textfield.render Mdl
                [ 2 ]
                model.mdl
                [ Textfield.label "Name (max.35 Characters)"
                , Textfield.floatingLabel
                , Textfield.text_
                , Options.onInput SetName
                , Textfield.value model.name
                ]
                []
            ]
        , div
            []
            [ Textfield.render Mdl
                [ 3 ]
                model.mdl
                [ Textfield.label "Symbol (3 or 4 characters)"
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
                [ Textfield.label "Description (max.285 Characters)"
                , Textfield.floatingLabel
                , Textfield.textarea
                , Textfield.rows 3
                , Options.onInput SetDescription
                , Textfield.value model.description
                ]
                []
            ]
        , renderHashtags ctx model
        ]


renderCategory active value =
    let
        background =
            case active == value of
                False ->
                    "#f2f2f2"

                True ->
                    "#1a8cf0"
    in
    Chip.button
        [ Options.css "margin" "0px 2px"
        , Options.css "background-color" background
        , Options.onClick (SetActiveCategory value)
        ]
        [ Chip.content []
            [ text value ]
        ]


renderHashtags ctx model =
    div []
        [ Textfield.render Mdl
            [ 5 ]
            model.mdl
            [ Textfield.label "Add a hashtag"
            , Textfield.floatingLabel
            , Textfield.text_
            , Textfield.value model.newHashtag
            , Options.onInput SetNewHashtag
            , Options.on "keydown" (JD.andThen isEnter keyCode)
            ]
            []
        , div [] <| List.indexedMap renderHashtag model.hashtags
        ]


renderHashtag index value =
    Chip.button
        [ Options.css "margin" "5px 5px"
        , Chip.deleteClick (RemoveHashtag index)
        , Chip.deleteIcon "cancel"
        ]
        [ Chip.content []
            [ text <| "# " ++ value ]
        ]


isEnter : number -> JD.Decoder Msg
isEnter code =
    if code == 13 then
        JD.succeed AddHashtag

    else
        JD.fail ""


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
