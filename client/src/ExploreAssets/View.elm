module ExploreAssets.View exposing (render, renderAsset)

import Common.Change24 as Change24
import Common.Styles
    exposing
        ( marginBottom
        , marginTop
        , paddingBottom
        , textCenter
        , textLeft
        , toMdlCss
        )
import ExploreAssets.Model exposing (Model)
import ExploreAssets.Msg exposing (Msg(..))
import Html exposing (..)
import Html.Attributes exposing (class, href, src, style)
import Identicon exposing (identicon)
import Main.Context exposing (Context)
import Main.Routing exposing (assetPath, createAssetPath)
import Material.List as Lists
import Material.Options exposing (css)
import Model.Asset exposing (Asset)


render : Context -> Model -> Html Msg
render ctx model =
    div []
        [ case List.length model.assets == 0 of
            True ->
                div [] [ text "There are no tokens yet!" ]

            False ->
                div [ marginBottom 10 ]
                    [ Lists.ul [] <| List.map renderAsset model.assets
                    ]
        ]


renderAsset : Asset -> Html Msg
renderAsset asset =
    let
        liStyle =
            toMdlCss <|
                style
                    [ ( "border-bottom", "1px solid #ddd" )
                    ]
    in
    Lists.li [ Lists.withBody, liStyle ]
        [ Lists.content []
            [ a [ href <| assetPath asset.id ]
                [ span
                    []
                    [ identicon "16px" asset.name
                    , text <| " " ++ asset.name
                    ]
                ]
            , Lists.body [ css "padding-top" "15px" ]
                [ text <| "Supply= " ++ toString asset.totalSupply ++ " " ++ asset.symbol
                , text <| " | Market Cap= 0FILS"
                ]
            ]
        , Lists.content2 []
            [ Lists.info2 [] [ text "24h" ]
            , Change24.render asset.change24
            ]
        ]
