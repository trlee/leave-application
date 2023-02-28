class LeaveType_Helpers {
    // Insert to datatable one row per element of data
    insertRows(id, data) {
        const target  = document.querySelector('#'+id)
        target.innerHTML = ''
        data.forEach(element => {
            let opt = document.createElement('tr')
            opt.id = element.ID 
            opt.innerHTML = `<td>${element.ID}</td>
                             <td>${element.leaveType}</td>
                             <td>${element.unpaid}</td>
                             <td>${element.limit}</td>
                             <td>${element.entitlementCalculation}</td>
                             <td>${element.gender}</td>
                             <td>${element.attachmentMandatory}</td>
                             <td>${element.encashmentLeave}</td>`
            target.appendChild(opt)
        });
    }

    // trigger datatable and row click event
    triggerDT(dt, dtBody){
        // trigger datatable
        console.log(dt, dtBody)
        const table = $('#'+dt).DataTable()

    }
}