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
@Table(name = "processing_status_type")
public class ProcessingStatusTypeEntity {

    @Id
    @Column(name = "processing_status_type")
    private String processingStatusTypeCd;
    @Column(name = "description")
    private String description;
    @Column(name = "active_bit")
    private Integer activeBit;

    @OneToMany(mappedBy = "processingStatusTypeEntity")
    private List<ActivityEntity>activityEntityList;

    @OneToMany(mappedBy = "processingStatusTypeEntity")
    private List<MessageEntity>messageEntityList;
}
