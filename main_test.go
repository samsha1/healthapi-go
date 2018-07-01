package main

import (
		"testing"
		"fmt"
		"net/http"
		"net/http/httptest"

)
var a App


func TestHost(t *testing.T){
	//Router := mux.NewRouter()
	//Router.HandleFunc("/host",a.host).Methods("GET")
	//a :=745
	 var domain string ="encoderslab.com"
	// ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	http.Error(w, "Failed", http.StatusServiceUnavailable)
	// }))

	expected,_ := http.NewRequest("GET","/host?hostname="+domain,nil)

	response := executeRequest(expected)

	// defer ts.Close()

	// if err != nil {
	// 	t.Errorf("expected no error got %v", err)
	// }
	rr := httptest.NewRecorder()
	//result:=a.host(expected)
	var expectedResult ="{\"results\":[{\"attrs\":{\"__name\":\"encoderslab.com\",\"acknowledgement\":0.0,\"acknowledgement_expiry\":0.0,\"action_url\":\"\",\"active\":true,\"address\":\"198.50.130.67\",\"address6\":\"\",\"check_attempt\":1.0,\"check_command\":\"hostalive\",\"check_interval\":120.0,\"check_period\":\"\",\"check_timeout\":null,\"command_endpoint\":\"\",\"display_name\":\"encoderslab.com\",\"downtime_depth\":0.0,\"enable_active_checks\":true,\"enable_event_handler\":true,\"enable_flapping\":false,\"enable_notifications\":true,\"enable_passive_checks\":true,\"enable_perfdata\":true,\"event_command\":\"\",\"flapping\":false,\"flapping_last_change\":1528797943.8422160149,\"flapping_negative\":242.0,\"flapping_positive\":0.0,\"flapping_threshold\":30.0,\"force_next_check\":false,\"force_next_notification\":false,\"groups\":[\"linux-servers\",\"mailservers\"],\"ha_mode\":0.0,\"icon_image\":\"\",\"icon_image_alt\":\"\",\"last_check\":1528797943.8421919346,\"last_check_result\":{\"active\":true,\"check_source\":\"health3.01cloud.com\",\"command\":[\"/usr/lib64/nagios/plugins/check_ping\",\"-H\",\"198.50.130.67\",\"-c\",\"5000,100%\",\"-w\",\"3000,80%\"],\"execution_end\":1528797943.8421390057,\"execution_start\":1528797939.7304470539,\"exit_status\":0.0,\"output\":\"PING OK - Packet loss = 0%, RTA = 92.83 ms\",\"performance_data\":[\"rta=92.834999ms;3000.000000;5000.000000;0.000000\",\"pl=0%;80;100;0\"],\"schedule_end\":1528797943.8421919346,\"schedule_start\":1528797939.7297410965,\"state\":0.0,\"type\":\"CheckResult\",\"vars_after\":{\"attempt\":1.0,\"reachable\":true,\"state\":0.0,\"state_type\":1.0},\"vars_before\":{\"attempt\":1.0,\"reachable\":true,\"state\":0.0,\"state_type\":1.0}},\"last_hard_state\":0.0,\"last_hard_state_change\":1528655641.8319671154,\"last_reachable\":true,\"last_state\":0.0,\"last_state_change\":1528655641.8319671154,\"last_state_down\":1528655520.1028110981,\"last_state_type\":1.0,\"last_state_unreachable\":0.0,\"last_state_up\":1528797943.842206955,\"max_check_attempts\":3.0,\"name\":\"encoderslab.com\",\"next_check\":1528798061.0622179508,\"notes\":\"\",\"notes_url\":\"\",\"original_attributes\":null,\"package\":\"_etc\",\"paused\":false,\"retry_interval\":60.0,\"severity\":8.0,\"source_location\":{\"first_column\":1.0,\"first_line\":11.0,\"last_column\":29.0,\"last_line\":11.0,\"path\":\"/etc/icinga2/zones.d/master/hosts/encoderslab.com/encoderslab.com.conf\"},\"state\":0.0,\"state_type\":1.0,\"templates\":[\"encoderslab.com\",\"generic-host\"],\"type\":\"Host\",\"vars\":{\"client_endpoint\":\"encoderslab.com\",\"disks\":{\"disk\":{}},\"http_vhosts\":{\"http\":{\"check_address\":\"localhost\",\"http_uri\":\"/\"}},\"notification\":{\"mail\":{\"groups\":[\"icingaadmins\"]}},\"os\":\"Linux\",\"server\":\"mailserver\"},\"version\":0.0,\"volatile\":false,\"zone\":\"master\"},\"joins\":{},\"meta\":{},\"name\":\"encoderslab.com\",\"type\":\"Host\"}]}"
	if result != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult,result)
	}
	// w := httptest.NewRecorder()

	// h := handler{
	// 	EndPointUrl: ts.URL,
	// }

	// h.Host(w, r)
	// if w.Code != http.StatusServiceUnavailable {
	// 	t.Fatalf("Expected %v but got %v", http.StatusServiceUnavailable, w.Code)
	// }

}

func TestAdd(t *testing.T){
	a = App{}
	result := a.add(3,4)
	if result !=7 {
		fmt.Printf("Expected %d but got %d",7,result)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
