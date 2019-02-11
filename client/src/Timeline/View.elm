module Timeline.View exposing (render)

import Common.Styles
    exposing
        ( background
        , color
        , padding
        , paddingLeft
        , paddingTop
        , textCenter
        , toMdlCss
        )
import Dict exposing (Dict)
import Gallery
import Gallery.Image as Image
import Html exposing (..)
import Html.Attributes exposing (class, height, href, property, src, style, width)
import Html.Events exposing (onClick)
import Identicon exposing (identicon)
import Json.Encode
import Main.Context exposing (Context)
import Main.Routing exposing (assetPath, profilePath)
import Material.Button as Button
import Material.Icon as Icon
import Material.List as Lists
import Material.Options as Options exposing (css)
import Material.Typography as Typography
import Timeline.Model exposing (Model, TimelineEntry, TimelineType(..))
import Timeline.Msg exposing (Msg(..))


render : Context -> Model -> Html Msg
render ctx model =
    div
        [ Common.Styles.paddingBottom 15 ]
        [ case model.data of
            Nothing ->
                div [] [ text "..." ]

            Just data ->
                case List.length data.entries < 1 of
                    True ->
                        Options.styled p
                            [ Typography.body1, toMdlCss (padding 50), toMdlCss textCenter ]
                            [ text "Timeline is empty" ]

                    False ->
                        div []
                            [ div [ entriesListStyle ] <|
                                List.map (renderEntry ctx model) data.entries
                            , case data.hasMore of
                                True ->
                                    renderLoadMore ctx model

                                False ->
                                    span [] []
                            ]
        ]


renderLoadMore : Context -> Model -> Html Msg
renderLoadMore ctx model =
    div [ loadMoreButoonStyle ]
        [ Button.render Mdl
            [ 1 ]
            model.mdl
            [ Button.raised
            , Button.ripple
            , Button.colored
            , Options.onClick LoadMore
            ]
            [ text "load more ..."
            ]
        ]


renderEntry : Context -> Model -> TimelineEntry -> Html Msg
renderEntry ctx model entry =
    let
        profileImageUrl =
            case entry.userProfileImageURL == "" of
                True ->
                    "static/images/default-user-icon.jpg"

                False ->
                    entry.userProfileImageURL

        showVerifyBtn =
            case ctx.user of
                Just user ->
                    case user.id == entry.oracleId of
                        True ->
                            entry.status == 0

                        False ->
                            False

                Nothing ->
                    False
    in
    div [ entryStyle ]
        [ div []
            [ Lists.li [ Lists.withSubtitle, css "padding-left" "8px" ]
                [ Lists.content []
                    [ Lists.avatarImage profileImageUrl [ css "border" "1px solid #ddd" ]
                    , text " "
                    , a [ href <| profilePath entry.userId ]
                        [ text <| entry.userName
                        ]
                    , case model.timelineType of
                        AssetTimeline _ ->
                            span [] []

                        _ ->
                            span []
                                [ text " posted in "
                                , a [ href (assetPath entry.assetId) ]
                                    [ text " "
                                    , text entry.assetName
                                    , text " "
                                    , identicon "15px" entry.assetName
                                    ]
                                ]
                    , Lists.subtitle [] [ text entry.createdAtHuman ]
                    ]
                ]
            ]
        , Options.styled p
            [ Typography.body1, toMdlCss (paddingLeft 10) ]
            [ text entry.text ]
        , case entry.ytVideoId == "" of
            True ->
                span [] []

            False ->
                iframe
                    [ width 560
                    , height 315
                    , src <| "https://www.youtube.com/embed/" ++ entry.ytVideoId
                    , property "frameborder" (Json.Encode.string "0")
                    , property "allowfullscreen" (Json.Encode.string "true")
                    ]
                    []
        , case List.length entry.images > 0 of
            False ->
                span [] []

            True ->
                div [ galleryContainerStyle ]
                    [ case Dict.get entry.blockId model.gallery of
                        Just g ->
                            Html.map (GalleryMsg entry.blockId) <|
                                Gallery.view imageConfig
                                    g
                                    [ Gallery.Arrows ]
                                    (imageSlides entry.images)

                        Nothing ->
                            span [] []
                    ]
        , case entry.status == 1 of
            True ->
                span [] []

            False ->
                div [ verifiedStatusStyle (entry.status == 1) ]
                    [ Options.styled span
                        [ Typography.caption ]
                        [ text "This  post needs to be approved by "
                        , a [ href (profilePath entry.oracleId) ] [ text entry.oracleName ]
                        , text ", the moderator of the topic"
                        ]
                    ]
        , case showVerifyBtn of
            True ->
                div [ paddingLeft 10 ]
                    [ Button.render Mdl
                        [ 1, -2 ]
                        model.mdl
                        [ Button.raised
                        , Button.ripple
                        , Button.colored
                        , Options.onClick (VerifyBlock True entry.blockId)
                        ]
                        [ text "accept post"
                        ]
                    , text " "
                    , Button.render Mdl
                        [ 2, -2 ]
                        model.mdl
                        [ Button.raised
                        , Button.ripple
                        , Options.onClick (VerifyBlock False entry.blockId)
                        ]
                        [ text "drop"
                        ]
                    ]

            False ->
                div [ toggleFavoriteStyle, onClick (ToggleFavorite entry.blockId) ]
                    [ case entry.didUserLike of
                        True ->
                            Icon.view "favorite"
                                [ Icon.size24
                                , toMdlCss <| color "red"
                                ]

                        False ->
                            Icon.view "favorite_border"
                                [ Icon.size24
                                ]
                    , span [] [ text <| toString entry.favoritesCount ]
                    ]
        , case entry.ethereumTransactionAddress == "" of
            True ->
                span [] []

            False ->
                div [ txLinkStyle ]
                    [ a
                        [ href <|
                            "https://rinkeby.etherscan.io/tx/"
                                ++ entry.ethereumTransactionAddress
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


entriesListStyle : Attribute a
entriesListStyle =
    style [ ( "text-align", "left" ) ]


entryStyle : Attribute a
entryStyle =
    style
        [ ( "padding-bottom", "10px" )
        , ( "margin-bottom", "10px" )
        , ( "border-bottom", "1px solid #ddd" )
        , ( "position", "relative" )
        ]


galleryContainerStyle : Attribute a
galleryContainerStyle =
    style
        []


loadMoreButoonStyle : Attribute a
loadMoreButoonStyle =
    style [ ( "text-align", "center" ) ]


toggleFavoriteStyle : Attribute a
toggleFavoriteStyle =
    style
        [ ( "padding", "15px 10px 0 10px" )
        , ( "display", "inline-block" )
        , ( "cursor", "pointer" )
        ]


verifiedStatusStyle : Bool -> Attribute a
verifiedStatusStyle isVerified =
    style
        [ ( "background", "white" )
        , ( "padding", "10px" )
        , ( "z-index", "3" )
        ]


imageConfig : Gallery.Config
imageConfig =
    Gallery.config
        { id = "image-gallery"
        , transition = 500
        , width = Gallery.pct 100
        , height = Gallery.px 400
        }


imageSlides : List String -> List ( String, Html msg )
imageSlides images =
    List.map (\x -> ( x, Image.slide x Image.Cover )) images
