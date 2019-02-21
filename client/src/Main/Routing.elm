module Main.Routing exposing
    ( Route(..)
    , assetPath
    , exploreAssetsPath
    , fastSignupPath
    , forgotPassPath
    , homepagePath
    , loginPath
    , matchers
    , parseLocation
    , portfolioPath
    , profilePath
    , signupPath
    , termsPath
    )

import Navigation exposing (Location)
import UrlParser exposing ((</>), Parser, int, map, oneOf, parseHash, s, string, top)


type Route
    = NotFoundRoute
    | HomepageRoute
    | ExploreAssetsRoute
    | PortfolioRoute
    | UserLoginRoute
    | UserForgotRoute
    | UserForgotResetPassRoute Int String
    | UserSignupRoute
    | UserFastSignupRoute
    | AssetRoute Int
    | ProfileRoute Int
    | CreateAssetRoute


matchers : Parser (Route -> a) a
matchers =
    oneOf
        [ map HomepageRoute top
        , map ExploreAssetsRoute (s "explore-assets")
        , map PortfolioRoute (s "portfolio")
        , map UserLoginRoute (s "login")
        , map UserSignupRoute (s "signup")
        , map UserFastSignupRoute (s "fsignup")
        , map UserForgotRoute (s "forgot")
        , map UserForgotResetPassRoute (s "forgot-reset-password" </> int </> string)
        , map AssetRoute (s "asset" </> int)
        , map ProfileRoute (s "profile" </> int)
        , map CreateAssetRoute (s "create-asset")
        ]


parseLocation : Location -> Route
parseLocation location =
    case parseHash matchers location of
        Just route ->
            route

        Nothing ->
            NotFoundRoute


homepagePath : String
homepagePath =
    "#"


assetPath : Int -> String
assetPath assetId =
    "#asset/" ++ toString assetId


exploreAssetsPath : String
exploreAssetsPath =
    "#explore-assets"


profilePath : Int -> String
profilePath userId =
    "#profile/" ++ toString userId


trendingPath : String
trendingPath =
    "#tokens"


portfolioPath : String
portfolioPath =
    "#portfolio"


loginPath : String
loginPath =
    "#login"


signupPath : String
signupPath =
    "#signup"


forgotPassPath : String
forgotPassPath =
    "#forgot"


fastSignupPath : String
fastSignupPath =
    "#fsignup"


termsPath : String
termsPath =
    "#site-terms"


privacyPath : String
privacyPath =
    "#site-privacy"
