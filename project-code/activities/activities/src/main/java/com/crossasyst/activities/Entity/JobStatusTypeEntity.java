package com.crossasyst.activities.Entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Table;

@Entity
@Data
@AllArgsConstructor
@NoArgsConstructor
@Table(name = "job_status")
public class JobStatusTypeEntity {
    @Column(name = "job_status_type_cd")
    private String jobStatusTypeCd;
    @Column(name = "description")
    private String description;
    @Column(name = "active_bit")
    private  Integer activeBit;

}
