class LeaveType_Helpers {
    // Insert to datatable one row per element of data
    insertRows(id, data) {
        console.log('I am here')
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
                             <td>${element.encashmentLeave}</td>
                             <td><a href="#" class="btn btn-danger" id="openUpdateLeave" data-bs-toggle="modal" data-bs-target="#updateLeave" onclick="return showUpdateValues(${element.ID},'${element.leaveType}',${element.unpaid},${element.limit},${element.entitlementCalculation},${element.gender},${element.attachmentMandatory},${element.encashmentLeave})"></i>Update</a></td>`
            target.appendChild(opt)
        });
    }

    // trigger datatable and row click event
    triggerDT(dt, dtBody){
        // trigger datatable
        const table = $('#'+dt).DataTable()

    }
}