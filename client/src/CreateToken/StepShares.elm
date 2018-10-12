module CreateToken.StepShares exposing (render)

import Common.Error exposing (renderHttpError)
import CreateToken.Msg exposing (Msg(..))
import Html exposing (..)
import Html.Attributes exposing (..)
import Material
import Material.Card as Card
import Material.Options as Options exposing (css)
import Material.Textfield as Textfield


render ctx model =
    div
        [ style [ ( "text-align", "center" ) ] ]
        [ Card.view
            [ css "width" "95px"
            , css "height" "95px"
            , css "background" "url('http://www.slate.com/content/dam/slate/articles/health_and_science/science/2017/06/170621_SCI_TreePlantingHoax.jpg.CROP.promo-xlarge2.jpg') center / cover"
            , css "margin" "15px auto"
            ]
            []
        , div
            [ style
                [ ( "30px", "0" )
                , ( "padding", "0 30px" )
                , ( "margin", "30px 0" )
                ]
            ]
            [ text "Define how many units of your token shall be created."
            ]
        , div []
            [ text "Total supply Quantity"
            ]
        , div [ style [ ( "15px", "0" ) ] ]
            [ Textfield.render Mdl
                [ 1 ]
                model.mdl
                [ css "width" "200px"
                , Textfield.label "e.g 10000"
                , Textfield.value model.shares
                , Options.onInput SetShares
                ]
                []
            ]
        , renderHttpError model.createTokenError
        ]
