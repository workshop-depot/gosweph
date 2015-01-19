package gosweph

/*
#include "swephexp.h"
*/
import "C"

//------------------------------------------------------------------------------

import "unsafe"

//------------------------------------------------------------------------------

func Swe_cs2degstr(t int32, a *string) *string {
	_t := C.centisec(t)
	var _a *C.char = C.CString(*a)
	defer C.free(unsafe.Pointer(_a))

	result := string(C.GoString(C.swe_cs2degstr(_t, _a)))
	return &result
}

func Swe_rise_trans(tjd_ut float64, ipl int32, starname *string, epheflag int32, rsmi int32, geopos *[3]float64, atpress float64, attemp float64, tret *float64, serr *string) int32 {
	_tjd_ut := C.double(tjd_ut)
	_ipl := C.int32(ipl)
	var _starname *C.char = C.CString(*starname)
	defer C.free(unsafe.Pointer(_starname))
	_epheflag := C.int32(epheflag)
	_rsmi := C.int32(rsmi)
	_geopos := (*C.double)(&geopos[0])
	_atpress := C.double(atpress)
	_attemp := C.double(attemp)
	_tret := C.double(*tret)
	var _serr *C.char = C.CString(*serr)
	defer C.free(unsafe.Pointer(_serr))

	result := int32(C.swe_rise_trans(_tjd_ut, _ipl, _starname, _epheflag, _rsmi, _geopos, _atpress, _attemp, (*C.double)(&_tret), _serr))
	*starname = string(C.GoString(_starname))
	*tret = float64(_tret)
	*serr = string(C.GoString(_serr))
	return result
}

func Swe_rise_trans_true_hor(tjd_ut float64, ipl int32, starname *string, epheflag int32, rsmi int32, geopos *[6]float64, atpress float64, attemp float64, horhgt float64, tret *float64, serr *string) int32 {
	_tjd_ut := C.double(tjd_ut)
	_ipl := C.int32(ipl)
	var _starname *C.char = C.CString(*starname)
	defer C.free(unsafe.Pointer(_starname))
	_epheflag := C.int32(epheflag)
	_rsmi := C.int32(rsmi)
	_geopos := (*C.double)(&geopos[0])
	_atpress := C.double(atpress)
	_attemp := C.double(attemp)
	_horhgt := C.double(horhgt)
	_tret := C.double(*tret)
	var _serr *C.char = C.CString(*serr)
	defer C.free(unsafe.Pointer(_serr))

	result := int32(C.swe_rise_trans_true_hor(_tjd_ut, _ipl, _starname, _epheflag, _rsmi, _geopos, _atpress, _attemp, _horhgt, (*C.double)(&_tret), _serr))
	*starname = string(C.GoString(_starname))
	*tret = float64(_tret)
	*serr = string(C.GoString(_serr))
	return result
}

func Swe_sidtime(tjd_ut float64) float64 {
	_tjd_ut := C.double(tjd_ut)

	result := float64(C.swe_sidtime(_tjd_ut))
	return result
}

func Swe_sidtime0(tjd_ut float64, eps float64, nut float64) float64 {
	_tjd_ut := C.double(tjd_ut)
	_eps := C.double(eps)
	_nut := C.double(nut)

	result := float64(C.swe_sidtime0(_tjd_ut, _eps, _nut))
	return result
}

func Swe_houses(tjd_ut float64, geolat float64, geolon float64, hsys int32, cusps *[6]float64, ascmc *[6]float64) int32 {
	_tjd_ut := C.double(tjd_ut)
	_geolat := C.double(geolat)
	_geolon := C.double(geolon)
	_hsys := C.int(hsys)
	_cusps := (*C.double)(&cusps[0])
	_ascmc := (*C.double)(&ascmc[0])

	result := int32(C.swe_houses(_tjd_ut, _geolat, _geolon, _hsys, _cusps, _ascmc))
	return result
}

func Swe_houses_armc(armc float64, geolat float64, eps float64, hsys int32, cusps *[6]float64, ascmc *[6]float64) int32 {
	_armc := C.double(armc)
	_geolat := C.double(geolat)
	_eps := C.double(eps)
	_hsys := C.int(hsys)
	_cusps := (*C.double)(&cusps[0])
	_ascmc := (*C.double)(&ascmc[0])

	result := int32(C.swe_houses_armc(_armc, _geolat, _eps, _hsys, _cusps, _ascmc))
	return result
}

func Swe_houses_ex(tjd_ut float64, iflag int32, geolat float64, geolon float64, hsys int32, cusps *[13]float64, ascmc *[10]float64) int32 {
	_tjd_ut := C.double(tjd_ut)
	_iflag := C.int32(iflag)
	_geolat := C.double(geolat)
	_geolon := C.double(geolon)
	_hsys := C.int(hsys)
	_cusps := (*C.double)(&cusps[0])
	_ascmc := (*C.double)(&ascmc[0])

	result := int32(C.swe_houses_ex(_tjd_ut, _iflag, _geolat, _geolon, _hsys, _cusps, _ascmc))
	return result
}

func Swe_house_pos(armc float64, geolat float64, eps float64, hsys int32, xpin *[6]float64, serr *string) float64 {
	_armc := C.double(armc)
	_geolat := C.double(geolat)
	_eps := C.double(eps)
	_hsys := C.int(hsys)
	_xpin := (*C.double)(&xpin[0])
	var _serr *C.char = C.CString(*serr)
	defer C.free(unsafe.Pointer(_serr))

	result := float64(C.swe_house_pos(_armc, _geolat, _eps, _hsys, _xpin, _serr))
	*serr = string(C.GoString(_serr))
	return result
}

func Swe_gauquelin_sector(tjd_ut float64, ipl int32, starname *string, iflag int32, imeth int32, geopos *[6]float64, atpress float64, attemp float64, dgsect *[6]float64, serr *string) float64 {
	_tjd_ut := C.double(tjd_ut)
	_ipl := C.int32(ipl)
	var _starname *C.char = C.CString(*starname)
	defer C.free(unsafe.Pointer(_starname))
	_iflag := C.int32(iflag)
	_imeth := C.int32(imeth)
	_geopos := (*C.double)(&geopos[0])
	_atpress := C.double(atpress)
	_attemp := C.double(attemp)
	_dgsect := (*C.double)(&dgsect[0])
	var _serr *C.char = C.CString(*serr)
	defer C.free(unsafe.Pointer(_serr))

	result := float64(C.swe_gauquelin_sector(_tjd_ut, _ipl, _starname, _iflag, _imeth, _geopos, _atpress, _attemp, _dgsect, _serr))
	*serr = string(C.GoString(_serr))
	return result
}

func Swe_cotrans(xpo *[6]float64, xpn *[6]float64, eps float64) {
	_xpo := (*C.double)(&xpo[0])
	_xpn := (*C.double)(&xpn[0])
	_eps := C.double(eps)

	C.swe_cotrans(_xpo, _xpn, _eps)
}

func Swe_cotrans_sp(xpo *[6]float64, xpn *[6]float64, eps float64) {
	_xpo := (*C.double)(&xpo[0])
	_xpn := (*C.double)(&xpn[0])
	_eps := C.double(eps)

	C.swe_cotrans_sp(_xpo, _xpn, _eps)
}

func Swe_get_planet_name(ipl int32, spname *string) {
	_ipl := C.int(ipl)
	var _spname *C.char = C.CString(*spname)
	defer C.free(unsafe.Pointer(_spname))

	C.swe_get_planet_name(_ipl, _spname)
	*spname = string(C.GoString(_spname))
}

func Swe_degnorm(x float64) float64 {
	_x := C.double(x)

	result := float64(C.swe_degnorm(_x))
	return result
}

func Swe_csnorm(p int32) int32 {
	_p := C.centisec(p)

	result := int32(C.swe_csnorm(_p))
	return result
}

func Swe_difcsn(p1 int32, p2 int32) int32 {
	_p1 := C.centisec(p1)
	_p2 := C.centisec(p2)

	result := int32(C.swe_difcsn(_p1, _p2))
	return result
}

func Swe_difdegn(p1 float64, p2 float64) float64 {
	_p1 := C.double(p1)
	_p2 := C.double(p2)

	result := float64(C.swe_difdegn(_p1, _p2))
	return result
}

func Swe_difcs2n(p1 int32, p2 int32) int32 {
	_p1 := C.centisec(p1)
	_p2 := C.centisec(p2)

	result := int32(C.swe_difcs2n(_p1, _p2))
	return result
}

func Swe_difdeg2n(p1 float64, p2 float64) float64 {
	_p1 := C.double(p1)
	_p2 := C.double(p2)

	result := float64(C.swe_difdeg2n(_p1, _p2))
	return result
}

func Swe_csroundsec(x int32) int32 {
	_x := C.centisec(x)

	result := int32(C.swe_csroundsec(_x))
	return result
}

func Swe_d2l(x float64) int64 {
	_x := C.double(x)

	result := int64(C.swe_d2l(_x))
	return result
}

func Swe_day_of_week(jd float64) int32 {
	_jd := C.double(jd)

	result := int32(C.swe_day_of_week(_jd))
	return result
}

/*
func Swe_cs2timestr(t int32, sep int32, suppressZero bool, a *string) *string {
	_t := C.int32(t)
	_sep := C.int32(sep)
	_suppressZero := C.bool(suppressZero)
	var _a *C.char = C.CString(*a)
	defer C.free(unsafe.Pointer(_a))

	result := *string(C.swe_cs2timestr(_t, _sep, _suppressZero, _a))
	return result
}
*/

/*
func Swe_cs2lonlatstr(t int32, pchar Char, mchar Char, s *string) *string {
	_t := C.int32(t)
	Char
	Char
	var _s *C.char = C.CString(*s)
	defer C.free(unsafe.Pointer(_s))

	result := *string(C.swe_cs2lonlatstr(_t, _pchar, _mchar, _s))
	return result
}
*/

func Swe_calc_ut(tjd_ut float64, ipl int32, iflag int32, xx *[6]float64, serr *string) int32 {
	_tjd_ut := C.double(tjd_ut)
	_ipl := C.int32(ipl)
	_iflag := C.int32(iflag)
	_xx := (*C.double)(&xx[0])
	var _serr *C.char = C.CString(*serr)
	defer C.free(unsafe.Pointer(_serr))

	result := int32(C.swe_calc_ut(_tjd_ut, _ipl, _iflag, _xx, _serr))
	*serr = string(C.GoString(_serr))
	return result
}

func Swe_calc(tjd_et float64, ipl int32, iflag int32, xx *[6]float64, serr *string) int32 {
	_tjd_et := C.double(tjd_et)
	_ipl := C.int(ipl)
	_iflag := C.int32(iflag)
	_xx := (*C.double)(&xx[0])
	var _serr *C.char = C.CString(*serr)
	defer C.free(unsafe.Pointer(_serr))

	result := int32(C.swe_calc(_tjd_et, _ipl, _iflag, _xx, _serr))
	*serr = string(C.GoString(_serr))
	return result
}

/*
func Swe_fixstar_ut(star *string, tjd_ut float64, iflag int64, xx *[6]float64, serr *string) int64 {
	var _star *C.char = C.CString(*star)
	defer C.free(unsafe.Pointer(_star))
	_tjd_ut := C.double(tjd_ut)
	_iflag := C.int64(iflag)
	_xx := (*C.double)(&xx[0])
	var _serr *C.char = C.CString(*serr)
	defer C.free(unsafe.Pointer(_serr))

	result := int64(C.swe_fixstar_ut(_star, _tjd_ut, _iflag, _xx, _serr))
	*star = string(C.GoString(_star))
	*serr = string(C.GoString(_serr))
	return result
}
*/

/*
func Swe_fixstar(star *string, tjd_et float64, iflag int64, xx *[6]float64, serr *string) int64 {
	var _star *C.char = C.CString(*star)
	defer C.free(unsafe.Pointer(_star))
	_tjd_et := C.double(tjd_et)
	_iflag := C.int64(iflag)
	_xx := (*C.double)(&xx[0])
	var _serr *C.char = C.CString(*serr)
	defer C.free(unsafe.Pointer(_serr))

	result := int64(C.swe_fixstar(_star, _tjd_et, _iflag, _xx, _serr))
	*star = string(C.GoString(_star))
	*serr = string(C.GoString(_serr))
	return result
}
*/

func Swe_set_topo(geolon float64, geolat float64, altitude float64) {
	_geolon := C.double(geolon)
	_geolat := C.double(geolat)
	_altitude := C.double(altitude)

	C.swe_set_topo(_geolon, _geolat, _altitude)
}

func Swe_set_sid_mode(sid_mode int32, t0 float64, ayan_t0 float64) {
	_sid_mode := C.int32(sid_mode)
	_t0 := C.double(t0)
	_ayan_t0 := C.double(ayan_t0)

	C.swe_set_sid_mode(_sid_mode, _t0, _ayan_t0)
}

func Swe_get_ayanamsa_ut(tjd_ut float64) float64 {
	_tjd_ut := C.double(tjd_ut)

	result := float64(C.swe_get_ayanamsa_ut(_tjd_ut))
	return result
}

func Swe_get_ayanamsa(tjd_et float64) float64 {
	_tjd_et := C.double(tjd_et)

	result := float64(C.swe_get_ayanamsa(_tjd_et))
	return result
}

func Swe_deltat(tjd float64) float64 {
	_tjd := C.double(tjd)

	result := float64(C.swe_deltat(_tjd))
	return result
}

func Swe_date_conversion(y int32, m int32, d int32, hour float64, c rune, tjd *float64) int32 {
	_y := C.int(y)
	_m := C.int(m)
	_d := C.int(d)
	_hour := C.double(hour)
	_c := C.char(c)
	_tjd := C.double(*tjd)

	result := int32(C.swe_date_conversion(_y, _m, _d, _hour, _c, (*C.double)(&_tjd)))
	*tjd = float64(_tjd)
	return result
}

func Swe_julday(year int32, month int32, day int32, hour float64, gregflag int32) float64 {
	_year := C.int(year)
	_month := C.int(month)
	_day := C.int(day)
	_hour := C.double(hour)
	_gregflag := C.int(gregflag)

	result := float64(C.swe_julday(_year, _month, _day, _hour, _gregflag))
	return result
}

func Swe_revjul(tjd float64, gregflag int32, year *int32, month *int32, day *int32, hour *float64) {
	_tjd := C.double(tjd)
	_gregflag := C.int(gregflag)
	_year := C.int(*year)
	_month := C.int(*month)
	_day := C.int(*day)
	_hour := C.double(*hour)

	C.swe_revjul(_tjd, _gregflag, (*C.int)(&_year), (*C.int)(&_month), (*C.int)(&_day), (*C.double)(&_hour))
	*year = int32(_year)
	*month = int32(_month)
	*day = int32(_day)
	*hour = float64(_hour)
}

func Swe_utc_to_jd(iyear int32, imonth int32, iday int32, ihour int32, imin int32, dsec float64, gregflag int32, dret *[6]float64, serr *string) {
	_iyear := C.int32(iyear)
	_imonth := C.int32(imonth)
	_iday := C.int32(iday)
	_ihour := C.int32(ihour)
	_imin := C.int32(imin)
	_dsec := C.double(dsec)
	_gregflag := C.int32(gregflag)
	_dret := (*C.double)(&dret[0])
	var _serr *C.char = C.CString(*serr)
	defer C.free(unsafe.Pointer(_serr))

	C.swe_utc_to_jd(_iyear, _imonth, _iday, _ihour, _imin, _dsec, _gregflag, _dret, _serr)
	*serr = string(C.GoString(_serr))
}

func Swe_jdet_to_utc(tjd_et float64, gregflag int32, iyear *int32, imonth *int32, iday *int32, ihour *int32, imin *int32, dsec *float64) {
	_tjd_et := C.double(tjd_et)
	_gregflag := C.int32(gregflag)
	_iyear := C.int32(*iyear)
	_imonth := C.int32(*imonth)
	_iday := C.int32(*iday)
	_ihour := C.int32(*ihour)
	_imin := C.int32(*imin)
	_dsec := C.double(*dsec)

	C.swe_jdet_to_utc(_tjd_et, _gregflag, (*C.int32)(&_iyear), (*C.int32)(&_imonth), (*C.int32)(&_iday), (*C.int32)(&_ihour), (*C.int32)(&_imin), (*C.double)(&_dsec))
	*iyear = int32(_iyear)
	*imonth = int32(_imonth)
	*iday = int32(_iday)
	*ihour = int32(_ihour)
	*imin = int32(_imin)
	*dsec = float64(_dsec)
}

func Swe_jdut1_to_utc(tjd_ut float64, gregflag int32, iyear *int32, imonth *int32, iday *int32, ihour *int32, imin *int32, dsec *float64) {
	_tjd_ut := C.double(tjd_ut)
	_gregflag := C.int32(gregflag)
	_iyear := C.int32(*iyear)
	_imonth := C.int32(*imonth)
	_iday := C.int32(*iday)
	_ihour := C.int32(*ihour)
	_imin := C.int32(*imin)
	_dsec := C.double(*dsec)

	C.swe_jdut1_to_utc(_tjd_ut, _gregflag, (*C.int32)(&_iyear), (*C.int32)(&_imonth), (*C.int32)(&_iday), (*C.int32)(&_ihour), (*C.int32)(&_imin), (*C.double)(&_dsec))
	*iyear = int32(_iyear)
	*imonth = int32(_imonth)
	*iday = int32(_iday)
	*ihour = int32(_ihour)
	*imin = int32(_imin)
	*dsec = float64(_dsec)
}

func Swe_get_tid_acc() float64 {

	result := float64(C.swe_get_tid_acc())
	return result
}

func Swe_set_tid_acc(t_acc float64) {
	_t_acc := C.double(t_acc)

	C.swe_set_tid_acc(_t_acc)
}

func Swe_time_equ(tjd_et float64, e *float64, serr *string) int32 {
	_tjd_et := C.double(tjd_et)
	_e := C.double(*e)
	var _serr *C.char = C.CString(*serr)
	defer C.free(unsafe.Pointer(_serr))

	result := int32(C.swe_time_equ(_tjd_et, (*C.double)(&_e), _serr))
	*e = float64(_e)
	*serr = string(C.GoString(_serr))
	return result
}

func Swe_set_ephe_path(path *string) {
	var _path *C.char = C.CString(*path)
	defer C.free(unsafe.Pointer(_path))

	C.swe_set_ephe_path(_path)
}

func Swe_set_jpl_file(fname *string) {
	var _fname *C.char = C.CString(*fname)
	defer C.free(unsafe.Pointer(_fname))

	C.swe_set_jpl_file(_fname)
}

func Swe_close() {

	C.swe_close()
}
