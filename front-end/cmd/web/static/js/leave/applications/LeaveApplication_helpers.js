class LeaveApplication_Helpers{
    // Insert to datatable one row per element of data
    insertRows(id, data){
        const target = document.querySelector('#'+id)
        target.innerHTML = ''
        data.forEach(element => {
            let opt = document.createElement('tr')
            opt.id = element.ID
            opt.innerHTML = `<td>${element.ID}</td>
                             <td>${element.leaveTypeID}</td>
                             <td>${element.employeeID}</td>
                             <td>${element.leaveType}</td>
                             <td>${element.applyReason}</td>
                             <td>${element.approved}</td>
                             <td>${element.cancelled}</td>
                             <td>${element.rejectReason}</td>
                             <td>${element.cancelReason}</td>
                             <td>${element.approvedAt}</td>
                             <td>${element.approvedBy}</td>
                             <td>${element.leaveDate}</td>
                             <td>${element.attachments}</td>
                             <td><a href="#" class="btn btn-danger" id="openUpdateLeaveApp" data-bs-toggle="modal" data-bs-target="#updateLeaveApp" onclick="return showUpdate(${element.ID})"></i>Update</a></td>
                             <td><a href="#" class="btn btn-danger" id="deleteLeave" onclick="return deleteValues(${element.ID})"></i>Delete</a></td>`
            target.appendChild(opt)
        })
    }
    // trigger datatable and row click event
    triggerDT(dt, dtBody){
        // trigger datatable
        const table = $('#'+dt).DataTable()
        
    }
}