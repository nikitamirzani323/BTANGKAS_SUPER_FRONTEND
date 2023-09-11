<script>
    import Home from "../agenadminrule/Home.svelte";
    import Entry from "../agenadminrule/Entry.svelte";

    let listAdminrule = [];
    let listAgen = [];
    let sData = "";
    let record = "";
    let totalrecord = 0;
    let adminrule_idadmin = 0;
    let adminrule_idagen = "";
    let adminrule_nmagen = "";
    let adminrule_name = "";
    let adminrule_rule = "";
    export let table_header_font = "";
    export let table_body_font = "";
    let token = localStorage.getItem("token");
    let akses_page = false;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "AGENADMINRULE-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            akses_page = true;
            initAdminrule();
        }
    }
    async function initAdminrule() {
        listAdminrule = [];
        const res = await fetch("/api/agenadminrule", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({}),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            let record_listagen = json.listagen;
            let no = 0;
            if (record != null) {
                totalrecord = record.length;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listAdminrule = [
                        ...listAdminrule,
                        {
                            adminrule_no: no,
                            adminrule_idadmin: record[i]["agenadminrule_id"],
                            adminrule_idagen: record[i]["agenadminrule_idagen"],
                            adminrule_agen: record[i]["agenadminrule_nmagen"],
                            adminrule_name: record[i]["agenadminrule_name"],
                            adminrule_rule: record[i]["agenadminrule_rule"],
                            adminrule_create: record[i]["agenadminrule_create"],
                            adminrule_update: record[i]["agenadminrule_update"],
                        },
                    ];
                }
            }
            for (var i = 0; i < record_listagen.length; i++) {
                listAgen = [
                    ...listAgen,
                    {
                        masteragen_id: record_listagen[i]["masteragen_id"],
                        masteragen_nmagen: record_listagen[i]["masteragen_nmagen"],
                    },
                ];
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleBackHalaman = () => {
        adminrule_idadmin = "";
        adminrule_rule = "";
        sData = "";
        listAdminrule = [];
        listAgen = [];
        handleRefreshData("all");
    };
    const handleEditData = (e) => {
        adminrule_idadmin = e.detail.id;
        adminrule_idagen = e.detail.idagen;
        adminrule_nmagen = e.detail.nmagen;
        adminrule_name = e.detail.name;
        adminrule_rule = e.detail.rule;
        sData = "Edit";
    };
    const handleRefreshData = (e) => {
        listAdminrule = [];
        listAgen = [];
        totalrecord = 0;
        setTimeout(function () {
            initAdminrule();
        }, 500);
    };
    initapp();
</script>

{#if akses_page == true}
    {#if sData == "" && adminrule_idadmin == ""}
        <Home
            on:handleEditData={handleEditData}
            on:handleRefreshData={handleRefreshData}
            {token}
            {table_header_font}
            {table_body_font}
            {listAdminrule}
            {listAgen}
            {totalrecord}
        />
    {:else}
        <Entry
            on:handleBackHalaman={handleBackHalaman}
            {sData}
            {token}
            {table_header_font}
            {table_body_font}
            {adminrule_idadmin}
            {adminrule_idagen}
            {adminrule_nmagen}
            {adminrule_name}
            {adminrule_rule}
        />
    {/if}
{/if}
