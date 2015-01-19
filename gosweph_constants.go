package gosweph

const (
	SE_ASC    int32 = 0
	SE_MC     int32 = 1
	SE_ARMC   int32 = 2
	SE_VERTEX int32 = 3
	SE_EQUASC int32 = 4 /* "equatorial ascendant" */
	SE_COASC1 int32 = 5 /* "co-ascendant" (W. Koch) */
	SE_COASC2 int32 = 6 /* "co-ascendant" (M. Munkasey) */
	SE_POLASC int32 = 7 /* "polar ascendant" (M. Munkasey) */
	SE_NASCMC int32 = 8

	SE_JUL_CAL  int32 = 0
	SE_GREG_CAL int32 = 1

	SE_SIDM_FAGAN_BRADLEY   int32 = 0
	SE_SIDM_LAHIRI          int32 = 1
	SE_SIDM_DELUCE          int32 = 2
	SE_SIDM_RAMAN           int32 = 3
	SE_SIDM_USHASHASHI      int32 = 4
	SE_SIDM_KRISHNAMURTI    int32 = 5
	SE_SIDM_DJWHAL_KHUL     int32 = 6
	SE_SIDM_YUKTESHWAR      int32 = 7
	SE_SIDM_JN_BHASIN       int32 = 8
	SE_SIDM_BABYL_KUGLER1   int32 = 9
	SE_SIDM_BABYL_KUGLER2   int32 = 10
	SE_SIDM_BABYL_KUGLER3   int32 = 11
	SE_SIDM_BABYL_HUBER     int32 = 12
	SE_SIDM_BABYL_ETPSC     int32 = 13
	SE_SIDM_ALDEBARAN_15TAU int32 = 14
	SE_SIDM_HIPPARCHOS      int32 = 15
	SE_SIDM_SASSANIAN       int32 = 16
	SE_SIDM_GALCENT_0SAG    int32 = 17
	SE_SIDM_J2000           int32 = 18
	SE_SIDM_J1900           int32 = 19
	SE_SIDM_B1950           int32 = 20
	SE_SIDM_USER            int32 = 255
	SE_SIDBIT_ECL_T0        int32 = 256

	SEFLG_JPLEPH int64 = 1 // use JPL ephemeris
	SEFLG_SWIEPH int64 = 2 // use SWISSEPH ephemeris, default
	SEFLG_MOSEPH int64 = 4 // use Moshier ephemeris

	SEFLG_HELCTR  int64 = 8   // return heliocentric position
	SEFLG_TRUEPOS int64 = 16  // return true positions, not apparent
	SEFLG_J2000   int64 = 32  // no precession, i.e. give J2000 equinox
	SEFLG_NONUT   int64 = 64  // no nutation, i.e. mean equinox of date
	SEFLG_SPEED3  int64 = 128 // speed from 3 positions (do not use it, SEFLG_SPEED is faster and preciser.)

	SEFLG_SPEED      int64 = 256         // high precision speed (analyt. comp.)
	SEFLG_NOGDEFL    int64 = 512         // turn off gravitational deflection
	SEFLG_NOABERR    int64 = 1024        // turn off 'annual' aberration of light
	SEFLG_EQUATORIAL int64 = 2048        // equatorial positions are wanted
	SEFLG_XYZ        int64 = 4096        // cartesian, not polar, coordinates
	SEFLG_RADIANS    int64 = 8192        // coordinates in radians, not degrees
	SEFLG_BARYCTR    int64 = 16384       // barycentric positions
	SEFLG_TOPOCTR    int64 = (32 * 1024) // topocentric positions
	SEFLG_SIDEREAL   int64 = (64 * 1024) // sidereal positions

	/* planet numbers for the ipl parameter in swe_calc() */

	SE_ECL_NUT   int32 = -1
	SE_SUN       int32 = 0
	SE_MOON      int32 = 1
	SE_MERCURY   int32 = 2
	SE_VENUS     int32 = 3
	SE_MARS      int32 = 4
	SE_JUPITER   int32 = 5
	SE_SATURN    int32 = 6
	SE_URANUS    int32 = 7
	SE_NEPTUNE   int32 = 8
	SE_PLUTO     int32 = 9
	SE_MEAN_NODE int32 = 10
	SE_TRUE_NODE int32 = 11
	SE_MEAN_APOG int32 = 12
	SE_OSCU_APOG int32 = 13
	SE_EARTH     int32 = 14
	SE_CHIRON    int32 = 15
	SE_PHOLUS    int32 = 16
	SE_CERES     int32 = 17
	SE_PALLAS    int32 = 18
	SE_JUNO      int32 = 19
	SE_VESTA     int32 = 20
	SE_INTP_APOG int32 = 21
	SE_INTP_PERG int32 = 22

	SE_NPLANETS    int32 = 23
	SE_FICT_OFFSET int32 = 40
	SE_NFICT_ELEM  int32 = 15

	/* Hamburger or Uranian "planets" */

	SE_CUPIDO   int32 = 40
	SE_HADES    int32 = 41
	SE_ZEUS     int32 = 42
	SE_KRONOS   int32 = 43
	SE_APOLLON  int32 = 44
	SE_ADMETOS  int32 = 45
	SE_VULKANUS int32 = 46
	SE_POSEIDON int32 = 47

	/* other fictitious bodies */

	SE_ISIS              int32 = 48
	SE_NIBIRU            int32 = 49
	SE_HARRINGTON        int32 = 50
	SE_NEPTUNE_LEVERRIER int32 = 51
	SE_NEPTUNE_ADAMS     int32 = 52
	SE_PLUTO_LOWELL      int32 = 53
	SE_PLUTO_PICKERING   int32 = 54

	SE_AST_OFFSET int32 = 10000

	EX_HOUSE_SYS_P int32 = 'P' //Placidus
	EX_HOUSE_SYS_K int32 = 'K' //Koch
	EX_HOUSE_SYS_O int32 = 'O' //Porphyrius
	EX_HOUSE_SYS_R int32 = 'R' //Regiomontanus
	EX_HOUSE_SYS_C int32 = 'C' //Campanus
	EX_HOUSE_SYS_A int32 = 'A' //or ‘E’	Equal (cusp 1 is Ascendant)
	EX_HOUSE_SYS_V int32 = 'V' //Vehlow equal (Asc. in middle of house 1)
	EX_HOUSE_SYS_W int32 = 'W' //Whole sign
	EX_HOUSE_SYS_X int32 = 'X' //axial rotation system
	EX_HOUSE_SYS_H int32 = 'H' //azimuthal or horizontal system
	EX_HOUSE_SYS_T int32 = 'T' //Polich/Page (“topocentric” system)
	EX_HOUSE_SYS_B int32 = 'B' //Alcabitus
	EX_HOUSE_SYS_M int32 = 'M' //Morinus
	EX_HOUSE_SYS_U int32 = 'U' //Krusinski-Pisa
	EX_HOUSE_SYS_G int32 = 'G' //Gauquelin sectors

	// appended; uncategorized

	// for swe_rise_transit() and swe_rise_transit_true_hor()
	SE_CALC_RISE int32 = 1
	SE_CALC_SET  int32 = 2

	// upper meridian transit (southern for northern geo. latitudes)
	SE_CALC_MTRANSIT int32 = 4

	// lower meridian transit (northern, below the horizon)
	SE_CALC_ITRANSIT int32 = 8

	/* the following bits can be added (or’ed) to SE_CALC_RISE or SE_CALC_SET */

	// for rising or setting of disc center
	SE_BIT_DISC_CENTER int32 = 256

	// for rising or setting of lower limb of disc
	SE_BIT_DISC_BOTTOM int32 = 8192

	// if refraction is not to be considered
	SE_BIT_NO_REFRACTION int32 = 512

	// in order to calculate civil twilight
	SE_BIT_CIVIL_TWILIGHT int32 = 1024

	// in order to calculate nautical twilight
	SE_BIT_NAUTIC_TWILIGHT int32 = 2048

	// in order to calculate astronomical twilight
	SE_BIT_ASTRO_TWILIGHT int32 = 4096

	// neglect the effect of distance on disc size
	SE_BIT_FIXED_DISC_SIZE int32 = (16 * 1024)
)
