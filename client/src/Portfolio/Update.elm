module Portfolio.Update exposing (update)

import Debug
import Main.Context exposing (Context)
import Material
import Model.Portfolio exposing (portfolioDecoder)
import Portfolio.Command exposing (loadPortfolioCmd)
import Portfolio.Model exposing (Model)
import Portfolio.Msg exposing (Msg(..))


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        OnLoadPortfolioResponse resp ->
            case resp of
                Ok portfolio ->
                    { model | portfolio = Just portfolio, error = Nothing } ! []

                Err error ->
                    Debug.log (toString error)
                        { model
                            | portfolio = Nothing
                            , error = Just "Error loading your portfolio, please try\n                        again!"
                        }
                        ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
