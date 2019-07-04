module ExploreAssets.View exposing (render, renderAsset)

import Asset.Model exposing (Asset)
import Common.Styles exposing (padding, paddingLeft, textCenter, toMdlCss)
import ExploreAssets.Model exposing (Model)
import ExploreAssets.Msg exposing (Msg(..))
import Html exposing (..)
import Html.Attributes exposing (class, href, src, style)
import Html.Events exposing (onClick)
import Identicon exposing (identicon)
import Main.Context exposing (Context)
import Main.Routing exposing (assetPath, profilePath)
import Material.Button as Button
import Material.Icon as Icon
import Material.List as Lists
import Material.Options as Options exposing (css)
import Material.Typography as Typography


render : Context -> Model -> Html Msg
render ctx model =
    div [ cardsWrapStyle ]
        [ case List.length model.assets == 0 of
            True ->
                div [] [ text "There are no tokens yet!" ]

            False ->
                div [ entriesListStyle ]
                    [ header [ textCenter ]
                        [ text "Tokens"
                        ]
                    , div [] <| List.map (renderAsset ctx model) model.assets
                    ]
        ]


renderAsset : Context -> Model -> Asset -> Html Msg
renderAsset ctx model asset =
    div [ entryStyle ]
        [ div []
            [ Lists.li [ Lists.withSubtitle, css "padding-left" "8px" ]
                [ Lists.content []
                    [ identicon "16px" asset.name
                    , text " "
                    , a [ href <| assetPath asset.id ]
                        [ text <| asset.name ]
                    , Options.styled span [ Typography.body1 ] [ text <| " - " ++ asset.symbol ]
                    , Lists.subtitle []
                        [ text <| "Supply= " ++ toString asset.totalSupply
                        , text <| " | Contributors: " ++ toString asset.minersCount
                        , text <| " | Oracle: "
                        , a [ href (profilePath asset.creatorId) ] [ text asset.creatorName ]
                        ]
                    ]
                ]
            ]
        , Options.styled p
            [ Typography.body1, toMdlCss (paddingLeft 10) ]
            [ text asset.description ]
        , div [ toggleFavoriteStyle, onClick (ToggleFavorite asset.id) ]
            [ case asset.didUserLike of
                True ->
                    Icon.view "favorite"
                        [ Icon.size24
                        , css "color" "red"
                        ]

                False ->
                    Icon.view "favorite_border"
                        [ Icon.size24
                        ]
            , span [] [ text <| toString asset.favoritesCount ]
            ]
        , div [ txLinkStyle ]
            [ a
                [ href <|
                    "https://rinkeby.etherscan.io/address/"
                        ++ asset.ethereumAddress
                , Html.Attributes.target "_blank"
                ]
                [ img [ txIconStyle, src "static/images/ethereum.png" ] [] ]
            ]
        ]


txLinkStyle : Attribute a
txLinkStyle =
    style
        [ ( "display", "inline-block" )
        , ( "float", "right" )
        , ( "padding", "15px 5px" )
        ]


txIconStyle : Attribute a
txIconStyle =
    style
        [ ( "display", "inline-block" )
        , ( "width", "30px" )
        , ( "hight", "30px" )
        ]


cardsWrapStyle : Attribute a
cardsWrapStyle =
    style
        []


entriesListStyle : Attribute a
entriesListStyle =
    style [ ( "text-align", "left" ) ]


entryStyle : Attribute a
entryStyle =
    style
        [ ( "margin", "10px 5px" )
        , ( "border", "1px solid #ddd" )
        , ( "border-radius", "8px" )
        , ( "position", "relative" )
        , ( "padding-bottom", "10px" )
        ]


toggleFavoriteStyle : Attribute a
toggleFavoriteStyle =
    style
        [ ( "padding", "15px 10px 0 10px" )
        , ( "display", "inline-block" )
        , ( "cursor", "pointer" )
        ]
