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
@Table(name = "data_channel")
public class DataChannelEntity {
    @Column(name = "data_channel_cd")
    private String dataChannelCd;
    @Column(name = "description")
    private String description;
    @Column(name = "active_bit")
    private  Integer activeBit;
}
