<script>
    import { Input } from "sveltestrap";
    
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
    export let listCurr = [];
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "COMPANY"
	let title_idcompany = ""
	let title_minbet = ""
    let sData = "";
    let sDataListBet = "";
    let sDataConfPoint = "";
    let myModal_newentry = "";
    let flag_code = false;
    let flag_btnsave = true;
    let name_field = "";
    let idcurr_field = "";
    let url_field = "";
    let owner_field = "";
    let phone_field = "";
    let email_field = "";
    let status_field = "";
    let create_field = "";
    let update_field = "";
    let idrecord = "";
    let searchHome = "";
    let filterHome = [];

    //COMPANY LISTBET
    let listbet_master = [];
    let listbet_company = [];
    let idrecord_listbet_field = 0;
    let idcompany_listbet_field = "";
    let minbet_listbet_field = 0;

     //COMPANY CONF POINT
    let listconfpoint_company = [];
    let idrecord_confpoint_field = 0;
    let idbet_confpoint_field = 0;
    let idcompany_confpoint_field = "";
    let point_confpoint_field = 0;


    let css_loader = "display: none;";
    let msgloader = "";
    
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_id
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_name
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
    
    const NewData = (e,id,name,idcurr,url,owner,email,phone,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            if(status == "N"){
                flag_btnsave = false;
            }else{
                flag_btnsave = true;
            }
            flag_code = true;
            idrecord = id
            name_field = name;
            idcurr_field = idcurr;
            url_field = url;
            owner_field = owner;
            phone_field = phone;
            email_field = email;
            status_field = status;
            create_field = create;
            update_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const call_listconfpoint = (e,f) => {
        sDataConfPoint = "New"
        title_idcompany = idcompany_listbet_field
        title_minbet = f
        idbet_confpoint_field = e
        idcompany_confpoint_field = idcompany_listbet_field
        point_confpoint_field = 0;
        call_listconfpoint_bycompany(e)
        myModal_newentry = new bootstrap.Modal(document.getElementById("modallistconfpoint"));
        myModal_newentry.show();
        
    };
    
    const call_listbet = (e) => {
        sDataListBet = "New"
        title_idcompany = e
        idrecord_listbet_field = 0;
        idcompany_listbet_field = e;
        minbet_listbet_field = 0;
        call_listbet_master(e)
        call_listbet_bycompany(e)
        myModal_newentry = new bootstrap.Modal(document.getElementById("modallistbet"));
        myModal_newentry.show();
        
    };
    const call_formlistbet = () => {
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalcrudlistbetcompany"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(idrecord.length == 4){
                flag = false
                msg += "The ID is maxlengt 4\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(owner_field == ""){
                flag = false
                msg += "The Owner is required\n"
            }
            if(phone_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(owner_field == ""){
                flag = false
                msg += "The Owner is required\n"
            }
            if(phone_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/companysave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CURR-SAVE",
                    company_id: idrecord.toUpperCase(),
                    company_name: name_field,
                    company_idcurr:idcurr_field,
                    company_nmowner: owner_field,
                    company_phoneowner: phone_field,
                    company_emailowner: email_field,
                    company_url: url_field,
                    company_status: status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function call_listconfpoint_bycompany(idbet) {
        listconfpoint_company = [];
        const res = await fetch("/api/companyconfpoint", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company_idbet: parseInt(idbet),
                company_id: idcompany_listbet_field,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listconfpoint_company = [
                        ...listconfpoint_company,
                        {
                        companyconf_no: no,
                        companyconf_id: record[i]["companyconf_id"],
                        companyconf_idbet: record[i]["companyconf_idbet"],
                        companyconf_idpoin: record[i]["companyconf_idpoin"],
                        companyconf_nmpoin: record[i]["companyconf_nmpoin"],
                        companyconf_poin: record[i]["companyconf_poin"],
                        companyconf_create: record[i]["companyconf_create"],
                        companyconf_update: record[i]["companyconf_update"],
                        },
                    ];
                }
            }
        }
    }
    async function call_listbet_master() {
        listbet_master = [];
        const res = await fetch("/api/listbet", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
               
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listbet_master = [
                        ...listbet_master,
                        {
                        lisbet_id: record[i]["lisbet_id"],
                        lisbet_minbet: record[i]["lisbet_minbet"],
                        },
                    ];
                }
            }
        }
    }
    async function call_listbet_bycompany(idcompany) {
        listbet_company = [];
        const res = await fetch("/api/companylistbet", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company_id: idcompany,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listbet_company = [
                        ...listbet_company,
                        {
                        companylistbet_no: no,
                        companylistbet_id: record[i]["companylistbet_id"],
                        companylistbet_minbet: record[i]["companylistbet_minbet"],
                        companylistbet_create: record[i]["companylistbet_create"],
                        companylistbet_update: record[i]["companylistbet_update"],
                        },
                    ];
                }
            }
        }
    }
    async function handleSave_listbet() {
        let flag = true
        let msg = ""
        let min_bet = value_listbetmaster(minbet_listbet_field)
        if(sDataListBet == "New"){
            if(idcompany_listbet_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(minbet_listbet_field == ""){
                flag = false
                msg += "The Minbet is required\n"
            }
            if(parseInt(min_bet) <= 0){
                flag = false
                msg += "The Minbet cannot 0\n"
            }
        }else{
            if(idrecord_listbet_field == "" || idrecord_listbet_field == 0){
                flag = false
                msg += "The ID is required\n"
            }
            if(idcompany_listbet_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(minbet_listbet_field == ""){
                flag = false
                msg += "The Minbet is required\n"
            }
            if(parseInt(min_bet) <= 0){
                flag = false
                msg += "The Minbet cannot 0\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/companylistbetsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataListBet,
                    page:"CURR-SAVE",
                    companylistbet_id: parseInt(idrecord_listbet_field),
                    companylistbet_idcompany: idcompany_listbet_field,
                    companylistbet_minbet:parseInt(min_bet),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField_listbet()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleSave_confpoint_generate() {
        let flag = true
        let msg = ""
        if(sDataConfPoint == "New"){
            if(idcompany_confpoint_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(idbet_confpoint_field == ""){
                flag = false
                msg += "The IDBET is required\n"
            }
        }else{
            if(idcompany_confpoint_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(idbet_confpoint_field == ""){
                flag = false
                msg += "The IDBET is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/companyconfpointsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataListBet,
                    page:"CURR-SAVE",
                    companyconfpoint_id: parseInt(idrecord_confpoint_field),
                    companyconfpoint_idbet: parseInt(idbet_confpoint_field),
                    Companyconfpoint_idcompany: idcompany_confpoint_field,
                    Companyconfpoint_point:parseInt(point_confpoint_field),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField_listbet()
                }
                msgloader = json.message;
                call_listconfpoint_bycompany(idbet_confpoint_field)
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    function clearField(){
        flag_code = false
        idrecord = "";
        name_field = "";
        idcurr_field = "";
        url_field = "";
        owner_field = "";
        phone_field = "";
        email_field = "";
        status_field = "";
        create_field = "";
        update_field = "";
    }
    function clearField_listbet(){
        flag_code = false
        idrecord_listbet_field = 0;
        idcompany_listbet_field = "";
        minbet_listbet_field = 0;
    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterTafsirMimpi = [];
                listHome = [];
                const tafsir = {
                    searchTafsirMimpi,
                };
                dispatch("handleTafsirMimpi", tafsir);
        }  
    };
    function uperCase(element) {
        function onInput(event) {
            element.value = element.value.toUpperCase();
            element.value = element.value.replace(/[^A-Z0-9]/gi, '');
        }
        element.addEventListener("input", onInput);
        return {
            destroy() {
                element.removeEventListener("input", onInput);
            },
        };
    }
    function status(e){
        let result = "DEACTIVE"
        if(e == "Y"){
            result = "ACTIVE"
        }
        return result
    }
    function value_listbetmaster(e){
        let value = 0
        for(let i=0;i<listbet_master.length;i++){
            if(e == listbet_master[i].lisbet_id){
                value = parseInt(listbet_master[i].lisbet_minbet)
            }
        }
        return value;
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container-fluid" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <Button
                on:click={callFunction}
                button_function="NEW"
                button_title="New"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchHome}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Code"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan=2>&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">&nbsp;</th>
                                <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CODE</th>
                                <th NOWRAP width="2%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CURR</th>
                                <th NOWRAP width="7%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">START</th>
                                <th NOWRAP width="7%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">END</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">COMPANY</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">OWNER</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">EMAIL</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                //e,id,name,idcurr,url,owner,email,phone,status,create,update
                                                NewData("Edit",rec.home_id, 
                                                rec.home_name, rec.home_idcurr,rec.home_url,
                                                rec.home_nmowner, rec.home_emailowner, 
                                                rec.home_phoneowner,rec.home_status,
                                                rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                call_listbet(rec.home_id)
                                            }} class="bi bi-file-earmark"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_idcurr}</td>
                                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_start}</td>
                                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_end}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        <a href="{rec.home_url}" target="_blank">{rec.home_name}</a>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nmowner}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_phoneowner}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_emailowner}</td>
                                </tr>
                            {/each}
                        </tbody>
                        {:else}
                        <tbody>
                            <tr>
                                <td colspan="20">
                                    <center>
                                        <Loader />
                                    </center>
                                </td>
                            </tr>
                        </tbody>
                        {/if} 
                    </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>

<Modal
	modal_id="modalentrycrud"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">CODE</label>
                    <input bind:value={idrecord}
                        use:uperCase  
                        disabled={flag_code}
                        class="required form-control"
                        maxlength="9"
                        type="text"
                        placeholder="CODE"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Default Currency</label>
                    <select
                        bind:value="{idcurr_field}" 
                        name="currency" id="currency" 
                        class="required form-control">
                        {#each listCurr as rec}
                        <option value="{rec.curr_id}">{rec.curr_id}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Company</label>
                    <Input bind:value={name_field}
                        class="required"
                        type="text"
                        placeholder="Company"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">URL</label>
                    <Input bind:value={url_field}
                        class="required"
                        type="text"
                        placeholder="URL"/>
                </div>
            </div>
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Owner</label>
                    <Input bind:value={owner_field}
                        class="required"
                        type="text"
                        placeholder="Owner"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone</label>
                    <Input bind:value={phone_field}
                        class="required"
                        type="text"
                        placeholder="Phone"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Email</label>
                    <Input bind:value={email_field}
                        class=""
                        type="text"
                        placeholder="Email"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={status_field}>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                    <div class="mb-3">
                        <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                            Create : {create_field}<br />
                            Update : {update_field}
                        </div>
                    </div>
                {/if}
            </div>
        </div>
        
        
        
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modallistbet"
	modal_size="modal-dialog-centered"
	modal_title="ListBet - {title_idcompany}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="*" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">MINBET</th>
                </tr>
            </thead>
            <tbody>
                {#each listbet_company as rec }
                    <tr>
                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                            <i on:click={() => {
                                    //e,id,name,idcurr,url,owner,email,phone,status,create,update
                                    call_listconfpoint(rec.companylistbet_id,rec.companylistbet_minbet);
                                }} class="bi bi-pencil"></i>
                        </td>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.companylistbet_no}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};">{new Intl.NumberFormat().format(rec.companylistbet_minbet)}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button on:click={() => {
                call_formlistbet();
            }} 
            button_title="New"
            button_css="btn-info"/>
	</slot:template>
</Modal>

<Modal
	modal_id="modalcrudlistbetcompany"
	modal_size="modal-dialog-centered"
	modal_title="{title_page+" "+title_idcompany+" / "+sDataListBet}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Min Bet</label>
            <select
                class="form-control required"
                style="text-align: right;"
                bind:value={minbet_listbet_field}>
                {#each listbet_master as rec}
                 <option value="{rec.lisbet_id}">{new Intl.NumberFormat().format(rec.lisbet_minbet)}</option>
                {/each}
            </select>
        </div>
        {#if sDataListBet != "New"}
            <div class="mb-3">
                <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                    Create : {create_field}<br />
                    Update : {update_field}
                </div>
            </div>
        {/if}
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave_listbet();
            }} 
            
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modallistconfpoint"
	modal_size="modal-dialog-centered"
	modal_title="Configure Point - {title_idcompany} - {title_minbet}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">POINT</th>
                </tr>
            </thead>
            <tbody>
                {#each listconfpoint_company as rec }
                    <tr>
                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                            <i on:click={() => {
                                    //e,id,name,idcurr,url,owner,email,phone,status,create,update
                                    NewData("Edit",rec.home_id, 
                                    rec.home_name, rec.home_idcurr,rec.home_url,
                                    rec.home_nmowner, rec.home_emailowner, 
                                    rec.home_phoneowner,rec.home_status,
                                    rec.home_create, rec.home_update);
                                }} class="bi bi-pencil"></i>
                        </td>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.companyconf_no}</td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyconf_nmpoin}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};">{new Intl.NumberFormat().format(rec.companyconf_poin)}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button on:click={() => {
                handleSave_confpoint_generate();
            }} 
            button_title="Generate"
            button_css="btn-info"/>
	</slot:template>
</Modal>