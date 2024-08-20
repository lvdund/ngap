from docx import Document
import re, os, sys, string
import datetime
import getopt
import getpass

msg_list = {}
type_list = {}

verbosity = 0
outdir = './'
cachedir = './cache/'
currentdir = './'

filename = "test.docx"
document = Document(filename)

FAIL = '\033[91m'
INFO = '\033[93m'
ENDC = '\033[0m'

def d_print(string):
    if verbosity > 0:
        sys.stdout.write(string)

def d_info(string):
    sys.stdout.write(INFO + string + ENDC + "\n")

def d_error(string):
    sys.stderr.write(FAIL + string + ENDC + "\n")
    sys.exit(0)

def v_upper(v):
    return re.sub('\'', '_', re.sub('/', '_', re.sub('-', '_', re.sub(' ', '_', v)))).upper()

def v_lower(v):
    return re.sub('\'', '_', re.sub('/', '_', re.sub('-', '_', re.sub(' ', '_', v)))).lower()

def get_cells(cells):
    iei = cells[0].text
    level = iei.count('>')
    presence = cells[1].text
    typeCell = re.sub(r'[0-9.]', '', cells[3].text)
    typeCell = typeCell.replace('\n', '')
    return { 
        "iei" : iei,
        "level": level,
        "presence" : presence,
        "type" : typeCell,
    }

def write_cells_to_file(name, cells):
    write_file(f, name + ".append({ \"iei\" : \"" + cells["iei"] + \
        "\", \"index\" : \"" + cells["index"] + \
        "\", \"level\" : \"" + cells["level"] + \
        "\", \"type\" : \"" + cells["type"] + \
        "\", \"presence\" : \"" + cells["presence"] + \
        "\", \"type\" : \"" + cells["type"] + "\"})\n")

def write_file(f, string):
    f.write(string)
    d_print(string)

def output_header_to_file(f):
    now = datetime.datetime.now()
    f.write("/**generated time: %s**/\n\n" % str(now))

#write data structure and methods for an IE
def write_ie_struct(f, ie, disabled):
    pass

#  write encode

# write decode





























# msg_list["PDU SESSION RESOURCE SETUP REQUEST"]["table"] = 4
# msg_list["PDU SESSION RESOURCE SETUP RESPONSE"]["table"] = 6
# msg_list["PDU SESSION RESOURCE RELEASE COMMAND"]["table"] = 8
# msg_list["PDU SESSION RESOURCE RELEASE RESPONSE"]["table"] = 10
# msg_list["PDU SESSION RESOURCE MODIFY REQUEST"]["table"] = 11
# msg_list["PDU SESSION RESOURCE MODIFY RESPONSE"]["table"] = 13
# msg_list["PDU SESSION RESOURCE NOTIFY"]["table"] = 15
# msg_list["PDU SESSION RESOURCE MODIFY INDICATION"]["table"] = 17
# msg_list["PDU SESSION RESOURCE MODIFY CONFIRM"]["table"] = 19
# msg_list["INITIAL CONTEXT SETUP REQUEST"]["table"] = 21
# msg_list["INITIAL CONTEXT SETUP RESPONSE"]["table"] = 24
# msg_list["INITIAL CONTEXT SETUP FAILURE"]["table"] = 26
# msg_list["UE CONTEXT RELEASE REQUEST"]["table"] = 28
# msg_list["UE CONTEXT RELEASE COMMAND"]["table"] = 30
# msg_list["UE CONTEXT RELEASE COMPLETE"]["table"] = 31
# msg_list["UE CONTEXT MODIFICATION REQUEST"]["table"] = 33
# msg_list["UE CONTEXT MODIFICATION RESPONSE"]["table"] = 34
# msg_list["UE CONTEXT MODIFICATION FAILURE"]["table"] = 35
# msg_list["RRC INACTIVE TRANSITION REPORT"]["table"] = 36
# msg_list["CONNECTION ESTABLISHMENT INDICATION"]["table"] = 37
# msg_list["AMF CP RELOCATION INDICATION"]["table"] = 38
# msg_list["RAN CP RELOCATION INDICATION"]["table"] = 39
# msg_list["RETRIEVE UE INFORMATION"]["table"] = 40
# msg_list["UE INFORMATION TRANSFER"]["table"] = 41
# msg_list["UE CONTEXT SUSPEND REQUEST"]["table"] = 42
# msg_list["UE CONTEXT SUSPEND RESPONSE"]["table"] = 44
# msg_list["UE CONTEXT SUSPEND FAILURE"]["table"] = 45
# msg_list["UE CONTEXT RESUME REQUEST"]["table"] = 46
# msg_list["UE CONTEXT RESUME RESPONSE"]["table"] = 48
# msg_list["UE CONTEXT RESUME FAILURE"]["table"] = 50
# msg_list["HANDOVER REQUIRED"]["table"] = 51
# msg_list["HANDOVER COMMAND"]["table"] = 53
# msg_list["HANDOVER PREPARATION FAILURE"]["table"] = 56
# msg_list["HANDOVER REQUEST"]["table"] = 57
# msg_list["HANDOVER REQUEST ACKNOWLEDGE"]["table"] = 59
# msg_list["HANDOVER FAILURE"]["table"] = 61
# msg_list["HANDOVER NOTIFY"]["table"] = 62
# msg_list["PATH SWITCH REQUEST"]["table"] = 63
# msg_list["PATH SWITCH REQUEST ACKNOWLEDGE"]["table"] = 65
# msg_list["PATH SWITCH REQUEST FAILURE"]["table"] = 67
# msg_list["HANDOVER CANCEL"]["table"] = 69
# msg_list["HANDOVER CANCEL ACKNOWLEDGE"]["table"] = 70
# msg_list["UPLINK RAN STATUS TRANSFER"]["table"] = 71
# msg_list["DOWNLINK RAN STATUS TRANSFER"]["table"] = 72
# msg_list["HANDOVER SUCCESS"]["table"] = 73
# msg_list["UPLINK RAN EARLY STATUS TRANSFER"]["table"] = 74
# msg_list["DOWNLINK RAN EARLY STATUS TRANSFER"]["table"] = 75
# msg_list["PAGING"]["table"] = 76
# msg_list["MULTICAST GROUP PAGING"]["table"] = 78
# msg_list["INITIAL UE MESSAGE"]["table"] = 80
# msg_list["DOWNLINK NAS TRANSPORT"]["table"] = 81
# msg_list["UPLINK NAS TRANSPORT"]["table"] = 82
# msg_list["NAS NON DELIVERY INDICATION"]["table"] = 83
# msg_list["REROUTE NAS REQUEST"]["table"] = 84
# msg_list["NG SETUP REQUEST"]["table"] = 85
# msg_list["NG SETUP RESPONSE"]["table"] = 87
# msg_list["NG SETUP FAILURE"]["table"] = 89
# msg_list["RAN CONFIGURATION UPDATE"]["table"] = 90
# msg_list["RAN CONFIGURATION UPDATE ACKNOWLEDGE"]["table"] = 92
# msg_list["RAN CONFIGURATION UPDATE FAILURE"]["table"] = 93
# msg_list["AMF CONFIGURATION UPDATE"]["table"] = 94
# msg_list["AMF CONFIGURATION UPDATE ACKNOWLEDGE"]["table"] = 96
# msg_list["AMF CONFIGURATION UPDATE FAILURE"]["table"] = 98
# msg_list["AMF STATUS INDICATION"]["table"] = 99
# msg_list["NG RESET"]["table"] = 101
# msg_list["NG RESET ACKNOWLEDGE"]["table"] = 102
# msg_list["ERROR INDICATION"]["table"] = 103
# msg_list["OVERLOAD START"]["table"] = 104
# msg_list["OVERLOAD STOP"]["table"] = 106
# msg_list["UPLINK RAN CONFIGURATION TRANSFER"]["table"] = 107
# msg_list["DOWNLINK RAN CONFIGURATION TRANSFER"]["table"] = 108
# msg_list["WRITE-REPLACE WARNING REQUEST"]["table"] = 109
# msg_list["WRITE-REPLACE WARNING RESPONSE"]["table"] = 110
# msg_list["PWS CANCEL REQUEST"]["table"] = 111
# msg_list["PWS CANCEL RESPONSE"]["table"] = 112
# msg_list["PWS RESTART INDICATION"]["table"] = 113
# msg_list["PWS FAILURE INDICATION"]["table"] = 115
# msg_list["DOWNLINK UE ASSOCIATED NRPPA TRANSPORT"]["table"] = 117
# msg_list["UPLINK UE ASSOCIATED NRPPA TRANSPORT"]["table"] = 118
# msg_list["DOWNLINK NON UE ASSOCIATED NRPPA TRANSPORT"]["table"] = 119
# msg_list["UPLINK NON UE ASSOCIATED NRPPA TRANSPORT"]["table"] = 120
# msg_list["TRACE START"]["table"] = 121
# msg_list["TRACE FAILURE INDICATION"]["table"] = 122
# msg_list["DEACTIVATE TRACE"]["table"] = 123
# msg_list["CELL TRAFFIC TRACE"]["table"] = 124
# msg_list["LOCATION REPORTING CONTROL"]["table"] = 125
# msg_list["LOCATION REPORTING FAILURE INDICATION"]["table"] = 127
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 
# msg_list[""]["table"] = 