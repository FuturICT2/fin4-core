module Timeline.Update exposing (update)

import Dict exposing (Dict)
import Gallery
import Main.Context exposing (Context)
import Material
import Timeline.Command exposing (loadTimelineCmd, toggleFavoriteCmd, verifyBlockCmd)
import Timeline.Model exposing (Model)
import Timeline.Msg exposing (Msg(..))


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    let
        page =
            case model.data of
                Just data ->
                    data.page

                Nothing ->
                    1
    in
    case msg of
        OnToggleFavoriteResponse resp ->
            model ! []

        ToggleFavorite blockId ->
            let
                updateDidUserLik item =
                    if item.blockId == blockId then
                        let
                            newDidUserLike =
                                not item.didUserLike

                            favoritesCount =
                                case newDidUserLike of
                                    True ->
                                        item.favoritesCount + 1

                                    False ->
                                        item.favoritesCount - 1
                        in
                        { item
                            | didUserLike = newDidUserLike
                            , favoritesCount = favoritesCount
                        }

                    else
                        item

                dataEntries =
                    case model.data of
                        Just d ->
                            d.entries

                        Nothing ->
                            []

                items =
                    List.map updateDidUserLik dataEntries

                newData =
                    case model.data of
                        Just data ->
                            Just { data | entries = items }

                        Nothing ->
                            model.data
            in
            { model | data = newData } ! [ toggleFavoriteCmd ctx blockId ]

        VerifyBlock isRejected blockId ->
            model ! [ verifyBlockCmd ctx isRejected blockId ]

        VerifyBlockResponse resp ->
            case resp of
                Ok data ->
                    model ! [ loadTimelineCmd ctx page model.timelineType ]

                Err error ->
                    model ! []

        LoadMore ->
            { model | isLoadingMore = True } ! [ loadTimelineCmd ctx (page + 1) model.timelineType ]

        OnLoadTimelineResponse resp ->
            case resp of
                Ok data ->
                    let
                        loadedDataEntreis =
                            case model.data of
                                Just d ->
                                    case data.page == 1 of
                                        True ->
                                            []

                                        False ->
                                            d.entries

                                Nothing ->
                                    []

                        newEntreis =
                            List.append loadedDataEntreis data.entries

                        newPage =
                            case List.length data.entries == 0 of
                                True ->
                                    data.page - 1

                                False ->
                                    data.page

                        newData =
                            { data | entries = newEntreis, page = newPage }

                        mapped =
                            List.map
                                (\v -> ( v.blockId, Gallery.init (List.length v.images) ))
                                newData.entries

                        gallery =
                            Dict.fromList mapped
                    in
                    { model
                        | data = Just newData
                        , gallery = gallery
                        , isLoadingMore = False
                        , loadMoreError = Nothing
                    }
                        ! []

                Err err ->
                    { model
                        | isLoadingMore = False
                        , loadMoreError = Just err
                    }
                        ! []

        GalleryMsg blockId msg_ ->
            let
                maybeGallery =
                    Dict.get blockId model.gallery

                prev =
                    Maybe.withDefault (Gallery.init 0) maybeGallery

                d =
                    Dict.insert blockId (Gallery.update msg_ prev) model.gallery
            in
            { model | gallery = d } ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
