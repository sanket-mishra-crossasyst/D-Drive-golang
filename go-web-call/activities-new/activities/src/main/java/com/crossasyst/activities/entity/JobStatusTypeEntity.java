package com.crossasyst.activities.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.util.List;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "job_status_type")
public class JobStatusTypeEntity {

    @Id
    @Column(name = "job_status_type")
    private String jobStatusType;
    @Column(name = "description")
    private String description;
    @Column(name = "active_bit")
    private Integer activeBit;

    @OneToMany(mappedBy = "jobStatusTypeEntity")
    private List<DataJobEntity>dataJobEntityList;
}
